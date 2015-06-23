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

func addOrdinaryWorkDay(workspaceId, projectId string, date string) {
	timezone, _ := time.LoadLocation("Europe/Rome")
	exampleFormat := "2006-01-02"
	midnightDate, errorDateFormat := time.ParseInLocation(exampleFormat, date, timezone)
	if errorDateFormat != nil {
		println("ERROR: Date format should be: 2015-05-31")
		println(errorDateFormat.Error())
		os.Exit(1)
	}

	morningStartTime := midnightDate.Add(9 * time.Hour)
	morningEntry := createHalfDayTimeEntry(workspaceId, projectId, morningStartTime, []string{})
	sendTimeEntry(morningEntry)

	afternoonStartTime := midnightDate.Add(14 * time.Hour)
	afternoonEntry := createHalfDayTimeEntry(workspaceId, projectId, afternoonStartTime, []string{})
	sendTimeEntry(afternoonEntry)
}

func addVacationDay(workspaceId, date string) {
	timezone, _ := time.LoadLocation("Europe/Rome")
	exampleFormat := "2006-01-02"
	midnightDate, errorDateFormat := time.ParseInLocation(exampleFormat, date, timezone)
	if errorDateFormat != nil {
		println("ERROR: Date format should be: 2015-05-31")
		println(errorDateFormat.Error())
		os.Exit(1)
	}

	vacationProjectId := "8352044"
	morningStartTime := midnightDate.Add(9 * time.Hour)
	morningEntry := createHalfDayTimeEntry(workspaceId, vacationProjectId, morningStartTime, []string{"Ferie"})
	sendTimeEntry(morningEntry)

	afternoonStartTime := midnightDate.Add(14 * time.Hour)
	afternoonEntry := createHalfDayTimeEntry(workspaceId, vacationProjectId, afternoonStartTime, []string{"Ferie"})
	sendTimeEntry(afternoonEntry)
}
