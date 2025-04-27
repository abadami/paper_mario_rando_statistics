package main

import (
	"fmt"
	"sync"
)

//TODO: Fail state for error'd task
//Essentially, anywhere in the process should error out and no data should be saved if possible
func FetchRaceDetailsFromRacetime() {
	fmt.Println("Fetching race details from racetime...")
	
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

	taskLogError := InsertTaskLog(true, len(results))

	if taskLogError != nil {
		fmt.Println("Error inserting task log error. Oh no!")
	}

	fmt.Println("Finished fetching race data from racetime!")
}
