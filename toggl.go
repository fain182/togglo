package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func sendTimeEntries(entries []TimeEntry) {
	for _, entry := range entries {
		sendTimeEntry(&entry)
	}
}

func sendTimeEntry(entry *TimeEntry) {
	entryJson, err := json.Marshal(entry)
	if err != nil {
		println("error: " + err.Error())
		return
	}
	var jsonStr = "{\"time_entry\":" + string(entryJson) + "}"
	sendApiRequest("POST", "/time_entries", jsonStr)
}

func getProjects() []Project {
	configuration := getConfiguration()
	response := sendApiRequest("GET", "/workspaces/"+configuration.WorkspaceId+"/projects", "")

	var projects []Project
	err := json.Unmarshal([]byte(response), &projects)
	if err != nil {
		fmt.Println("error:", err)
	}
	return projects
}

func sendApiRequest(method, togglRelativeUrl, jsonBody string) string {
	configuration := getConfiguration()
	url := "https://www.toggl.com/api/v8" + togglRelativeUrl
	req, err := http.NewRequest(method, url, bytes.NewBufferString(jsonBody))
	if err != nil {
		println("error: " + err.Error())
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configuration.ApiToken, "api_token")

	var client = http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		println("error: " + err.Error())
		return ""
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	if resp.StatusCode != 200 {
		println(resp.Status)
		println(body)
		os.Exit(1)
	}

	return body
}
