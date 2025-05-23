package main

import (
	"fmt"
	"math"
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

func CalculateAverage(times []int) int {
	sum := 0
	
	for _, seconds := range times {
		sum += seconds
	}

	average := sum / len(times)

	return average
}

func CalculateDeviation(times []int, average int, count int) float64 {
	deviationSum := 0.0

	for _, seconds := range times {
		deviationSum += math.Pow(float64(seconds - average), 2)
	}

	deviationAverage := deviationSum / float64(count)

	return math.Sqrt(float64(deviationAverage))
}