package env

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Load() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if len(value) > 1 && (strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") || strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			value = value[1 : len(value)-1]
		}

		if _, exists := os.LookupEnv(key); !exists {
			if err := os.Setenv(key, value); err != nil {
				log.Printf("Error setting env var %s: %v", key, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading .env file: %v", err)
	}
}

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return fallback
}
