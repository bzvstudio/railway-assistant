package types

import "time"

type RailwayStatus struct {
	ActiveIncidents []Incident `json:"activeIncidents"`
	RecentIncidents []Incident `json:"recentIncidents"`
}

type Incident struct {
	ID         string              `json:"id"`
	Slug       string              `json:"slug"`
	Title      string              `json:"title"`
	Status     string              `json:"status"`
	CreatedAt  time.Time           `json:"createdAt"`
	ResolvedAt *time.Time          `json:"resolvedAt"`
	Components []IncidentComponent `json:"components"`
	Updates    []IncidentUpdate    `json:"updates"`
}

type IncidentComponent struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
	Impact    string `json:"impact"`
}

type IncidentUpdate struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
