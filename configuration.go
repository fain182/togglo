package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

type Configuration struct {
	ApiToken    string
	WorkspaceId string
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
