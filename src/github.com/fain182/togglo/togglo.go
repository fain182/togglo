package main

import (
	"bytes"
	"encoding/json"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
)

type Configuration struct {
	ApiToken    string
	WorkspaceId string
}

type TimeEntry struct {
	Description string `json:"description"`
	CreatedWith string `json:"created_with"`
	Start       string `json:"start"`
	Duration    int    `json:"duration"`
	WorkspaceId string `json:"wid"`
	ProjectId   string `json:"pid"`
	Billable    bool   `json:"billable"`
}

func main() {
	app := cli.NewApp()
	app.Name = "Togglo"
	app.Usage = "Compile toggle from command-line"
	app.Commands = []cli.Command{
		{
			Name:  "work",
			Usage: "Add a ordinary work time entry",
			Action: func(c *cli.Context) {
				var projectId = c.Args()[0]
				var date = c.Args()[1]
				addOrdinaryWorkDay(projectId, date)
			},
		},
	}

	app.Run(os.Args)
}

func addOrdinaryWorkDay(projectId string, date string) {
	configuration := getConfiguration()
	var morningEntry = &TimeEntry{
		Description: "dev",
		CreatedWith: "Togglo",
		Start:       date + "T09:00:00+02:00",
		Duration:    14400,
		WorkspaceId: configuration.WorkspaceId,
		ProjectId:   projectId,
		Billable:    true,
	}
	var afternoonEntry = &TimeEntry{
		Description: "dev",
		CreatedWith: "Togglo",
		Start:       date + "T14:00:00+02:00",
		Duration:    14400,
		WorkspaceId: configuration.WorkspaceId,
		ProjectId:   projectId,
		Billable:    true,
	}

	sendTimeEntry(morningEntry)
	sendTimeEntry(afternoonEntry)
}

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
	println(resp.Status)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	println(s)
}

func getConfiguration() Configuration {
	usr, err := user.Current()
	if err != nil {
		println("error: " + err.Error())
	}
	file, err := ioutil.ReadFile(usr.HomeDir + "/.togglo.json")
	if err != nil {
		println("error:" + err.Error())
		os.Exit(1)
	}
	var conf Configuration
	err = json.Unmarshal(file, &conf)
	if err != nil {
		println("error:" + err.Error())
		os.Exit(1)
	}
	return conf
}
