package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Togglo"
	app.Usage = "Compile toggle from command-line"
	app.Commands = []cli.Command{
		{
			Name:  "work",
			Usage: "Add a ordinary work time entry",
			Action: func(c *cli.Context) {
				configuration := getConfiguration()
				var projectId = c.Args()[0]
				var date = c.Args()[1]
				addOrdinaryWorkDay(configuration.WorkspaceId, projectId, date)
			},
		},
	}

	app.Run(os.Args)
}

func addOrdinaryWorkDay(workspaceId, projectId string, date string) {
	if len(date) != len("2015-01-12") {
		println("ERROR: Date format should be: 2015-05-05")
		os.Exit(1)
	}
	morningEntry := createHalfDayTimeEntry(workspaceId, projectId, date+"T09:00:00+02:00")
	afternoonEntry := createHalfDayTimeEntry(workspaceId, projectId, date+"T14:00:00+02:00")
	sendTimeEntry(morningEntry)
	sendTimeEntry(afternoonEntry)
}
