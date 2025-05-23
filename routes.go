package main

import (
	"github.com/jackc/pgx/v5"
)

func GetRaceAverageByFilters(request StatisticsRequest) (StatisticsResponse, error) {
	entrantId, getEntrantError := GetEntrant(pgx.NamedArgs{
		"entrantName": request.ContainsEntrant,
	})
	
	if getEntrantError != nil {
		return StatisticsResponse{}, getEntrantError
	}

	results, error := GetRacesByRaceEntrant(pgx.NamedArgs{
		"entrantId": entrantId,
	})

	if error != nil {
		return StatisticsResponse{}, error
	}

	count := 0

	dnfCount := 0

	var times []int
	
	for _, raceDetail := range results {
		if (raceDetail.Status == "dnf") {
			dnfCount += 1
			continue
		}
		time := ParseTimeString(raceDetail.Finish_time)
		times = append(times, time)
	}

	average := CalculateAverage(times)

	deviation := CalculateDeviation(times, average, count)

	return StatisticsResponse{
		Average: ParseSecondsToTime(average),
		Deviation: ParseSecondsToTime(int(deviation)),
		RaceNumber: count,
		DnfCount: dnfCount,
		RawData: results, 
	}, nil
}