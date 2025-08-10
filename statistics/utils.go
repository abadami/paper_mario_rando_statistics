package statistics

import (
	"fmt"
	"math"
	"regexp"

	"github.com/abadami/randomizer-statistics/domain"
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

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds_tracker)
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
		deviationSum += math.Pow(float64(seconds-average), 2)
	}

	deviationAverage := deviationSum / float64(count)

	return math.Sqrt(float64(deviationAverage))
}

func CalculateBestWinAndAverageWin(racerData []domain.RaceEntrantAndRaceRecord, fullData []domain.RaceEntrantAndRaceRecord) (int, int) {
	bestWin := math.MaxInt
	sum := 0
	count := 0

	for _, entrantRace := range racerData {
		if entrantRace.Place != 1 {
			continue
		}

		filteredRaces := filterRacesByRaceId(fullData, entrantRace.Race_id)

		for _, race := range filteredRaces {
			if race.Id == entrantRace.Id || race.Status == "dnf" {
				continue
			}

			entrantRaceFinishTime := ParseTimeString(entrantRace.Finish_time)
			raceFinishTime := ParseTimeString(race.Finish_time)

			difference := raceFinishTime - entrantRaceFinishTime

			sum += difference
			count += 1

			if difference < bestWin {
				bestWin = difference
			}
		}
	}

	if count == 0 {
		return bestWin, 0
	}

	return bestWin, sum / count
}

func CalculateWorstLoss(racerData []domain.RaceEntrantAndRaceRecord, fullData []domain.RaceEntrantAndRaceRecord) int {
	worstLose := 0

	for _, entrantRace := range racerData {
		if entrantRace.Place == 1 {
			continue
		}

		filteredRaces := filterRacesByRaceId(fullData, entrantRace.Race_id)

		firstPlace := getFirstPlaceRaceData(filteredRaces)

		entrantRaceFinishTime := ParseTimeString(entrantRace.Finish_time)
		raceFinishTime := ParseTimeString(firstPlace.Finish_time)

		difference := entrantRaceFinishTime - raceFinishTime

		if difference > worstLose {
			worstLose = difference
		}
	}

	return worstLose
}

func filterRacesByRaceId(data []domain.RaceEntrantAndRaceRecord, id int) []domain.RaceEntrantAndRaceRecord {
	filteredRaces := []domain.RaceEntrantAndRaceRecord{}

	for _, race := range data {
		if race.Race_id == id && race.Status != "dnf" {
			filteredRaces = append(filteredRaces, race)
		}
	}

	return filteredRaces
}

func getFirstPlaceRaceData(data []domain.RaceEntrantAndRaceRecord) domain.RaceEntrantAndRaceRecord {
	defaultRace := domain.RaceEntrantAndRaceRecord{}

	for _, race := range data {
		if race.Place == 1 {
			return race
		}
	}

	return defaultRace
}
