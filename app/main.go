package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/abadami/randomizer-statistics/internal/repositories/postgres"
	"github.com/abadami/randomizer-statistics/internal/repositories/racetime"
	racetime_service "github.com/abadami/randomizer-statistics/racetime"
	"github.com/abadami/randomizer-statistics/rest"
	"github.com/abadami/randomizer-statistics/statistics"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/robfig/cron/v3"

	"github.com/joho/godotenv"
)

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool, err := postgres.CreatePool()

	if err != nil {
		fmt.Println("Oh no, failed to connect to db!")
		return
	}

	defer dbpool.Close()

	//Setup repos
	raceRepo := postgres.NewRaceRepository(dbpool)
	entrantRepo := postgres.NewEntrantRepository(dbpool)
	tasklogRepo := postgres.NewTaskLogRepository(dbpool)
	racetimeRepo := racetime.NewRacetimeRepository()

	//Setup services
	racetimeService := racetime_service.NewService(racetimeRepo, raceRepo, tasklogRepo)
	statisticsService := statistics.NewService(raceRepo)

	c := cron.New()
	c.AddFunc("0 * * * *", func() {
		fmt.Print("Job is running!")
		racetimeService.FetchRaceDetailsFromRacetime()
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
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "client/dist"))
	FileServer(r, "/", filesDir)

	//Setup handlers
	rest.NewEntrantHandler(r, entrantRepo)
	rest.NewStatisticsHandler(r, statisticsService)

	fmt.Println("Server is Up")

	http.ListenAndServe(":3000", r)
}
