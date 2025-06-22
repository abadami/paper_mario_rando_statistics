package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/robfig/cron/v3"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool, err := CreatePool()

	if (err != nil) {
		fmt.Println("Oh no, failed to connect to db!")
		return
	}

	defer dbpool.Close()

	c := cron.New()
	c.AddFunc("0 * * * *", func() {
		fmt.Print("Job is running!")
		FetchRaceDetailsFromRacetime()
	})
	c.Start()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Get("/api/get_statistics_for_entrant", func(w http.ResponseWriter, r *http.Request) {
		entrant := r.URL.Query().Get("ContainsEntrant") //chi.URLParam(r, "ContainsEntrant")

		entrant_id, conversionError := strconv.Atoi(entrant)

		if conversionError != nil {
			http.Error(w, "Not a valid id value for contains entrant", http.StatusBadRequest)
		}

		response, error := GetRaceAverageByFilters(StatisticsRequest{
			ContainsEntrant: entrant_id,
		})

		if error != nil {
			http.Error(w, "Failed to get resource", http.StatusInternalServerError)
		}

		render.JSON(w, r, response)
	})

	r.Get("/api/get_race_entrants", func(w http.ResponseWriter, r *http.Request) {
		response, error := GetRaceEntrants()

		if error != nil {
			http.Error(w, "Failed to get resources", http.StatusInternalServerError)
		}

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
