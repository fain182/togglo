package main

type TimeEntry struct {
	Description string `json:"description"`
	CreatedWith string `json:"created_with"`
	Start       string `json:"start"`
	Duration    int    `json:"duration"`
	WorkspaceId string `json:"wid"`
	ProjectId   string `json:"pid"`
	Billable    bool   `json:"billable"`
}

func createHalfDayTimeEntry(workspaceId, projectId, datetime string) *TimeEntry {
	return &TimeEntry{
		Description: "dev",
		CreatedWith: "Togglo",
		Start:       datetime,
		Duration:    14400,
		WorkspaceId: workspaceId,
		ProjectId:   projectId,
		Billable:    true,
	}

}
