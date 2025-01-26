package main

import (
	"time"
)

func ParseTimeString(str string) float64 {
	dur, err := time.ParseDuration(str)

	if err != nil {
		return 0
	}

	return dur.Seconds()
}
