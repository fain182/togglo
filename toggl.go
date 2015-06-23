package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

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
	sendApiRequest("/time_entries", jsonStr)
}

func sendApiRequest(togglRelativeUrl, jsonBody string) {
	configuration := getConfiguration()
	url := "https://www.toggl.com/api/v8" + togglRelativeUrl
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(jsonBody))
	if err != nil {
		println("error: " + err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configuration.ApiToken, "api_token")

	var client = http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		println("error: " + err.Error())
		return
	}

	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		errorBody := buf.String()
		println(resp.Status)
		println(errorBody)
		os.Exit(1)
	}
}
