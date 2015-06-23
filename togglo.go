package main

import (
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Togglo"
	app.Usage = "Compile toggle from command-line"
	configuration := getConfiguration()
	app.Commands = []cli.Command{
		{
			Name:  "work",
			Usage: "Add a ordinary work day time entry",
			Action: func(c *cli.Context) {
				var projectId = c.Args()[0]
				var date = c.Args()[1]
				addOrdinaryWorkDay(configuration.WorkspaceId, projectId, date)
			},
		},
		{
			Name:  "vacation",
			Usage: "Add a vacation day time entry",
			Action: func(c *cli.Context) {
				var date = c.Args()[0]
				addVacationDay(configuration.WorkspaceId, date)
			},
		},
	}

	app.Run(os.Args)
}

func addOrdinaryWorkDay(workspaceId, projectId, date string) {
	midnightDate := parseDate(date)
	entries := createDayTimeEntries(workspaceId, projectId, midnightDate, []string{})
	sendTimeEntries(entries)
}

func addVacationDay(workspaceId, date string) {
	midnightDate := parseDate(date)
	vacationProjectId := "8352044"
	entries := createDayTimeEntries(workspaceId, vacationProjectId, midnightDate, []string{"Ferie"})
	sendTimeEntries(entries)
}

func parseDate(date string) time.Time {
	timezone, _ := time.LoadLocation("Europe/Rome")
	exampleFormat := "2006-01-02"
	midnightDate, errorDateFormat := time.ParseInLocation(exampleFormat, date, timezone)
	if errorDateFormat != nil {
		println("ERROR: Date format should be: 2015-05-31")
		println(errorDateFormat.Error())
		os.Exit(1)
	}
	return midnightDate
}
