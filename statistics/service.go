package statistics

import (
	"github.com/abadami/randomizer-statistics/domain"
	"github.com/abadami/randomizer-statistics/internal/utils"
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
		time := utils.ParseTimeString(raceDetail.Finish_time)
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

	average := utils.CalculateAverage(times)

	deviation := utils.CalculateDeviation(times, average, count)

	bestWin, averageWin := utils.CalculateBestWinAndAverageWin(entrantData, results)

	worstLoss := utils.CalculateWorstLoss(entrantData, results)

	//TODO: Determine if we even need to return the "full raw data". Users can just get that from racetime lol
	return domain.StatisticsResponse{
		Average:     utils.ParseSecondsToTime(average),
		Deviation:   utils.ParseSecondsToTime(int(deviation)),
		BestWin:     utils.ParseSecondsToTime(bestWin),
		WorstLoss:   utils.ParseSecondsToTime(worstLoss),
		AverageWin:  utils.ParseSecondsToTime(averageWin),
		RaceNumber:  count,
		DnfCount:    dnfCount,
		RawData:     entrantData,
		FullRawData: results,
	}, nil
}
