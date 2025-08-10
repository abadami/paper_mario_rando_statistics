package rest

import (
	"net/http"
	"strconv"

	"github.com/abadami/randomizer-statistics/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type StatisticsService interface {
	GetStatisticsForEntrant(request domain.StatisticsRequest) (domain.StatisticsResponse, error)
}

type StatisticsHandler struct {
	statisticsService StatisticsService
}

// This has logic that should be in a handler
func NewStatisticsHandler(c *chi.Mux, service StatisticsService) {
	handler := &StatisticsHandler{
		statisticsService: service,
	}
	c.Get("/api/get_statistics_for_entrant", handler.GetRaceAverageByFilters)
}

func (handler *StatisticsHandler) GetRaceAverageByFilters(w http.ResponseWriter, r *http.Request) {
	entrant := r.URL.Query().Get("ContainsEntrant") //chi.URLParam(r, "ContainsEntrant")
	goal := r.URL.Query().Get("Goal")
	raceType := r.URL.Query().Get("RaceType")

	entrant_id := -1

	if entrant != "" {
		paramConversion, conversionError := strconv.Atoi(entrant)

		if conversionError != nil {
			http.Error(w, "Not a valid id value for contains entrant", http.StatusBadRequest)
		}

		entrant_id = paramConversion
	}

	response, error := handler.statisticsService.GetStatisticsForEntrant(domain.StatisticsRequest{
		ContainsEntrant: entrant_id,
		Goal:            goal,
		RaceType:        raceType,
	})

	if error != nil {
		http.Error(w, "Failed to get resource", http.StatusInternalServerError)
	}

	render.JSON(w, r, response)
}
