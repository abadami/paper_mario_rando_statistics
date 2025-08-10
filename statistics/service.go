package statistics

import (
	"github.com/abadami/randomizer-statistics/domain"
)

type RaceRepository interface {
	GetRacesByRaceEntrant(request domain.StatisticsRequest) ([]domain.RaceEntrantAndRaceRecord, error)
}

type StatisticsService struct {
	raceRepo RaceRepository
}

func NewService(raceRepo RaceRepository) *StatisticsService {
	return &StatisticsService{
		raceRepo: raceRepo,
	}
}

func (service *StatisticsService) GetStatisticsForEntrant(request domain.StatisticsRequest) (domain.StatisticsResponse, error) {
	results, error := service.raceRepo.GetRacesByRaceEntrant(request)

	if error != nil {
		return domain.StatisticsResponse{}, error
	}

	entrantData := []domain.RaceEntrantAndRaceRecord{}

	for _, raceDetail := range results {
		if raceDetail.Entrant_id == request.ContainsEntrant {
			entrantData = append(entrantData, raceDetail)
		}
	}

	count := 0

	dnfCount := 0

	data := entrantData

	//Need to handle the "all" data situation
	if request.ContainsEntrant == -1 {
		data = results
	}

	var times []int

	for _, raceDetail := range data {
		if raceDetail.Status == "dnf" {
			dnfCount += 1
			continue
		}
		count += 1
		time := ParseTimeString(raceDetail.Finish_time)
		times = append(times, time)
	}

	if count == 0 {
		return domain.StatisticsResponse{
			Average:     "00:00:00",
			Deviation:   "00:00:00",
			BestWin:     "00:00:00",
			WorstLoss:   "00:00:00",
			AverageWin:  "00:00:00",
			RaceNumber:  0,
			DnfCount:    0,
			RawData:     []domain.RaceEntrantAndRaceRecord{},
			FullRawData: []domain.RaceEntrantAndRaceRecord{},
		}, nil
	}

	average := CalculateAverage(times)

	deviation := CalculateDeviation(times, average, count)

	bestWin, averageWin := CalculateBestWinAndAverageWin(entrantData, results)

	worstLoss := CalculateWorstLoss(entrantData, results)

	//TODO: Determine if we even need to return the "full raw data". Users can just get that from racetime lol
	return domain.StatisticsResponse{
		Average:     ParseSecondsToTime(average),
		Deviation:   ParseSecondsToTime(int(deviation)),
		BestWin:     ParseSecondsToTime(bestWin),
		WorstLoss:   ParseSecondsToTime(worstLoss),
		AverageWin:  ParseSecondsToTime(averageWin),
		RaceNumber:  count,
		DnfCount:    dnfCount,
		RawData:     entrantData,
		FullRawData: results,
	}, nil
}
