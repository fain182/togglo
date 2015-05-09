package main

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	exampleConf := []byte(`{"ApiToken":"ABC","WorkspaceId":"123"}`)
	result := parseConfiguration(exampleConf)
	if result.ApiToken != "ABC" {
		t.Errorf("ApiToken is wrong: %s", result.ApiToken)
	}
	if result.WorkspaceId != "123" {
		t.Errorf("WorkspaceId is wrong: %s", result.WorkspaceId)
	}
}
