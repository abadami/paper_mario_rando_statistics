package main

import (
	"fmt"
	"regexp"

	"github.com/senseyeio/duration"
)

func ParseTimeString(str string) int {
	re := regexp.MustCompile(`\.[0-9]*`)
	
	filteredString := re.ReplaceAllString(str, "")

	dur, err := duration.ParseISO8601(filteredString)

	if err != nil {
		return 0
	}

	//Duration in seconds
	return dur.TS + (dur.TM * 60) + (dur.TH * 60 * 60)
}

func ParseSecondsToTime(seconds int) string {
	seconds_tracker := seconds

	hours := seconds_tracker / 3600

	seconds_tracker -= hours * 3600

	minutes := seconds_tracker / 60

	seconds_tracker -= minutes * 60

	return fmt.Sprintf("%d:%d:%d", hours, minutes, seconds_tracker)
}