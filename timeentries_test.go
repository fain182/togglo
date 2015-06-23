package main

import (
	"testing"
)

func TestMorningIsFrom9(t *testing.T) {
	entries := createDayTimeEntries("123", "456", parseDate("2015-03-04"), []string{})
	if entries[0].Start != "2015-03-04T09:00:00+01:00" {
		t.Errorf("Start time is wrong: %s", entries[0].Start)
	}
}

func TestAfternoonIsFrom14(t *testing.T) {
	entries := createDayTimeEntries("123", "456", parseDate("2015-03-04"), []string{})
	if entries[1].Start != "2015-03-04T14:00:00+01:00" {
		t.Errorf("Start time is wrong: %s", entries[1].Start)
	}
}
