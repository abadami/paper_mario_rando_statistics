package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FetchRaceDetailsFromRacetime() {
	fmt.Println("Fetching race details from racetime...")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/randomizer_stats", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PWD"), os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_PORT"))

	dbpool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	defer dbpool.Close()

	racesResponse := GetRaceTitlesAndEntrantsByPage(1)

	jobs := make(chan int, racesResponse.NumPages)
	detailJobs := make(chan string, racesResponse.Count)
	results := make(chan RaceDetail, racesResponse.Count)

	pagewg := new(sync.WaitGroup)
	racewg := new(sync.WaitGroup)

	for _, race := range racesResponse.Races {
		detailJobs <- race.Name
		racewg.Add(1)
	}

	pagewg.Add(racesResponse.NumPages - 1)

	for w := 0; w <= 10; w++ {
		go GetPageWorker(GetPageWorkerParams{
			Id:      w,
			Jobs:    jobs,
			Results: detailJobs,
			Pagewg:  pagewg,
			Racewg:  racewg,
		})
	}

	for page := 2; page <= racesResponse.NumPages; page++ {
		jobs <- page
	}
	close(jobs)

	for w := 0; w <= 1000; w++ {
		go GetRaceWorker(GetRaceWorkerParams{
			Id:      w,
			Jobs:    detailJobs,
			Results: results,
			Wg:      racewg,
		})
	}

	pagewg.Wait()

	//We know we have all of the detailJobs, so we can close here
	close(detailJobs)

	racewg.Wait()

	//We know we have all the results now, so close the channels
	close(results)

	sum := 0
	count := 0

	var times []int

	for raceDetail := range results {
		fmt.Println(raceDetail.Name, " ", raceDetail.Entrants[0].FinishTime)
		time := ParseTimeString(raceDetail.Entrants[0].FinishTime)
		times = append(times, time)
		sum += time
		count += 1
	}

	average := sum / count

	deviation := CalculateDeviation(times, average, count)
	fmt.Println("Average time:", average)
	fmt.Println("Deviation:", deviation)

	fmt.Println("Finished fetching race details")
}
