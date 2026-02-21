package types

import "time"

type RailwayAlert struct {
	Type      string       `json:"type"`
	Severity  string       `json:"severity"`
	Timestamp time.Time    `json:"timestamp"`
	Resource  Resource     `json:"resource"`
	Details   AlertDetails `json:"details"`
}

type AlertDetails struct {
	ID            string `json:"id"`
	Branch        string `json:"branch"`
	Source        string `json:"source"`
	Status        string `json:"status"`
	CommitHash    string `json:"commitHash"`
	RepoSource    string `json:"repoSource"`
	CommitAuthor  string `json:"commitAuthor"`
	CommitMessage string `json:"commitMessage"`
}

type Resource struct {
	Workspace   Workspace   `json:"workspace"`
	Project     Project     `json:"project"`
	Environment Environment `json:"environment"`
	Service     *Service    `json:"service,omitempty"`
	Deployment  *Deployment `json:"deployment,omitempty"`
}

type Workspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Environment struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsEphemeral bool   `json:"isEphemeral"`
}

type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Deployment struct {
	ID string `json:"id"`
}
