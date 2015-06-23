package main

import "time"

type TimeEntry struct {
	Description string   `json:"description"`
	CreatedWith string   `json:"created_with"`
	Start       string   `json:"start"`
	Duration    int      `json:"duration"`
	WorkspaceId string   `json:"wid"`
	ProjectId   string   `json:"pid"`
	Billable    bool     `json:"billable"`
	Tags        []string `json:"tags"`
}

func createHalfDayTimeEntry(workspaceId, projectId string, datetime time.Time, tags []string) *TimeEntry {
	return &TimeEntry{
		Description: "dev",
		CreatedWith: "Togglo",
		Start:       datetime.Format(time.RFC3339),
		Duration:    14400,
		WorkspaceId: workspaceId,
		ProjectId:   projectId,
		Billable:    true,
		Tags:        tags,
	}

}
