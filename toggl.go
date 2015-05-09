package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sendTimeEntry(entry *TimeEntry) {
	configuration := getConfiguration()
	entryJson, err := json.Marshal(entry)
	var jsonStr = "{\"time_entry\":" + string(entryJson) + "}"
	req, err := http.NewRequest("POST", "https://www.toggl.com/api/v8/time_entries", bytes.NewBufferString(jsonStr))
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
		println(errorBody)
	}
}
