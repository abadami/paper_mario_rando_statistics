package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abadami/randomizer-statistics/domain"
	"github.com/abadami/randomizer-statistics/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type RaceRepository interface {
	GetRacesByRaceEntrant(request domain.StatisticsRequest) ([]domain.RaceEntrantAndRaceRecord, error)
}

type StatisticsHandler struct {
	raceRepository RaceRepository
}

//This has logic that should be in a handler
func NewStatisticsHandler(c *chi.Mux, repo RaceRepository) {
	handler := &StatisticsHandler{
		raceRepository: repo,
	}
	c.Get("/api/get_statistics_for_entrant", func(w http.ResponseWriter, r *http.Request) {
		entrant := r.URL.Query().Get("ContainsEntrant") //chi.URLParam(r, "ContainsEntrant")
		goal := r.URL.Query().Get("Goal")
		raceType := r.URL.Query().Get("raceType")

		entrant_id := -1

		if (entrant != "") {
			paramConversion, conversionError := strconv.Atoi(entrant)

			if conversionError != nil {
				http.Error(w, "Not a valid id value for contains entrant", http.StatusBadRequest)
			}

			entrant_id = paramConversion
		}

		response, error := handler.GetRaceAverageByFilters(domain.StatisticsRequest{
			ContainsEntrant: entrant_id,
			Goal: goal,
			RaceType: raceType,
		})

		if error != nil {
			http.Error(w, "Failed to get resource", http.StatusInternalServerError)
		}

		render.JSON(w, r, response)
	})
}

func (handler *StatisticsHandler) GetRaceAverageByFilters(request domain.StatisticsRequest) (domain.StatisticsResponse, error) {
fmt.Printf("%d", request.ContainsEntrant)
	results, error := handler.raceRepository.GetRacesByRaceEntrant(request)

	if error != nil {
		return domain.StatisticsResponse{}, error
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
		time := utils.ParseTimeString(raceDetail.Finish_time)
		times = append(times, time)
	}

	if count == 0 {
		return domain.StatisticsResponse{
			Average: "00:00:00",
			Deviation: "00:00:00",
			RaceNumber: 0,
			DnfCount: 0,
			RawData: []domain.RaceEntrantAndRaceRecord{}, 
		}, nil
	}

	average := utils.CalculateAverage(times)

	deviation := utils.CalculateDeviation(times, average, count)

	return domain.StatisticsResponse{
		Average: utils.ParseSecondsToTime(average),
		Deviation: utils.ParseSecondsToTime(int(deviation)),
		RaceNumber: count,
		DnfCount: dnfCount,
		RawData: results, 
	}, nil
}