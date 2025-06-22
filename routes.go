package main

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetRaceAverageByFilters(request StatisticsRequest) (StatisticsResponse, error) {
	fmt.Printf("%d", request.ContainsEntrant)
	results, error := GetRacesByRaceEntrant(pgx.NamedArgs{
		"entrantId": request.ContainsEntrant,
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
		count += 1
		time := ParseTimeString(raceDetail.Finish_time)
		times = append(times, time)
	}

	if count == 0 {
		return StatisticsResponse{
			Average: "00:00:00",
			Deviation: "00:00:00",
			RaceNumber: 0,
			DnfCount: 0,
			RawData: []RaceEntrantAndRaceRecord{}, 
		}, nil
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

func GetRaceEntrants() ([]EntrantRecord, error) {
	entrants, err := GetEntrants()

	if err != nil {
		return nil, err
	}

	return entrants, nil
}