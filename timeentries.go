package main

import "time"

type TimeEntry struct {
	Description string `json:"description"`
	CreatedWith string `json:"created_with"`
	Start       string `json:"start"`
	Duration    int    `json:"duration"`
	WorkspaceId string `json:"wid"`
	ProjectId   string `json:"pid"`
	Billable    bool   `json:"billable"`
}

func createHalfDayTimeEntry(workspaceId, projectId string, datetime time.Time) *TimeEntry {
	return &TimeEntry{
		Description: "dev",
		CreatedWith: "Togglo",
		Start:       datetime.Format(time.RFC3339),
		Duration:    14400,
		WorkspaceId: workspaceId,
		ProjectId:   projectId,
		Billable:    true,
	}

}
