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
	file, err := ioutil.ReadFile(getConfigurationFilePath())
	if err != nil {
		println("error:" + err.Error())
		os.Exit(1)
	}
	return parseConfiguration(file)
}

func getConfigurationFilePath() string {
	usr, err := user.Current()
	if err != nil {
		println("error: " + err.Error())
	}
	return usr.HomeDir + "/.togglo.json"
}

func parseConfiguration(file []byte) Configuration {
	var conf Configuration
	err := json.Unmarshal(file, &conf)
	if err != nil {
		println("error:" + err.Error())
		os.Exit(1)
	}
	return conf
}
