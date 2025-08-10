package rest

import (
	"net/http"

	"github.com/abadami/randomizer-statistics/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type EntrantRepository interface {
	GetEntrants() ([]domain.EntrantRecord, error)
}

type EntrantHandler struct {
	entrantRepository EntrantRepository
}

func NewEntrantHandler(c *chi.Mux, repo EntrantRepository) {
	handler := &EntrantHandler{
		entrantRepository: repo,
	}
	c.Get("/api/get_race_entrants", func(w http.ResponseWriter, r *http.Request) {
		response, error := handler.GetRaceEntrants()

		if error != nil {
			http.Error(w, "Failed to get resources", http.StatusInternalServerError)
		}

		render.JSON(w, r, response)
	})
}

func (handler *EntrantHandler) GetRaceEntrants() ([]domain.EntrantRecord, error) {
	entrants, err := handler.entrantRepository.GetEntrants()

	if err != nil {
		return nil, err
	}

	return entrants, nil
}