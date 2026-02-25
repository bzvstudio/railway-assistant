package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"strings"
	"time"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			b := make([]byte, 16)
			rand.Read(b)
			requestID = hex.EncodeToString(b)
		}
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "requestID", requestID)))
	})
}

func RealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			i := strings.Index(xff, ", ")
			if i == -1 {
				i = len(xff)
			}
			r.RemoteAddr = xff[:i]
		} else if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
			r.RemoteAddr = xrip
		}
		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &responseWriter{ResponseWriter: w, code: http.StatusOK}

		defer func() {
			if r.URL.Path == "/health" && ww.code == http.StatusOK {
				return
			}

			if isNoise(r.URL.Path) {
				return
			}

			log.Printf(
				"[%s] %s %s %d %s",
				r.Method,
				r.URL.Path,
				r.RemoteAddr,
				ww.code,
				time.Since(start),
			)
		}()

		next.ServeHTTP(ww, r)
	})
}

func isNoise(path string) bool {
	exts := []string{
		".php", ".env", ".git", ".yaml", ".xml",
		".css", ".js", ".map", ".png", ".jpg", ".ico",
		".asp", ".aspx", ".jsp", ".cgi",
		".aws", ".s3cfg", ".conf", ".config",
	}

	for _, ext := range exts {
		if strings.Contains(path, ext) {
			return true
		}
	}

	prefixes := []string{
		"/wp-", "/actuator", "/_debugbar", "/debug",
		"/api/", "/console", "/shell",
		"/wordpress", "/wlwmanifest",
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}

	return false
}

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("PANIC: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
