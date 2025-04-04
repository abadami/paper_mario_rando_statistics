package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/robfig/cron/v3"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := cron.New()
	c.AddFunc("0 * * * *", func() {
		fmt.Print("Job is running!")
		FetchRaceDetailsFromRacetime()
	})
	c.Start()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		entrant := chi.URLParam(r, "ContainsEntrant")

		response := GetRaceAverageByFilters(StatisticsRequest{
			ContainsEntrant: entrant,
		})

		render.JSON(w, r, response)
	})

	r.Get("/api/run_race_job", func(w http.ResponseWriter, r *http.Request) {
		FetchRaceDetailsFromRacetime()

		response := struct {
			success bool
		}{
			success: true,
		}

		render.JSON(w, r, response)
	})
	http.ListenAndServe(":3000", r)
}
