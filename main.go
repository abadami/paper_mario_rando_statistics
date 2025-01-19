package main

import (
	"fmt"
	"sync"
)

func getPageWorker(id int, jobs <-chan int, results chan<- string, pagewg *sync.WaitGroup, racewg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("getPageWorker", id, "started job", j)
		response := getRaceTitlesAndEntrantsByPage(j)
		for race := 0; race < len(response.Races); race++ {
			fmt.Println("Processing race title ", response.Races[race].Name)
			racewg.Add(1)
			results <- response.Races[race].Name
		}
		pagewg.Done()
	}
}

func getRaceWorker(id int, jobs <-chan string, results chan<- RaceDetail, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("getRaceWorker", id, "started job", j)
		response := getRaceDetails(j)
		fmt.Println("Processed race ", response.Name)
		results <- response
		wg.Done()
	}
}

func main() {
	racesResponse := getRaceTitlesAndEntrantsByPage(1)

	jobs := make(chan int, racesResponse.NumPages)
	detailJobs := make(chan string, racesResponse.Count)
	results := make(chan RaceDetail, racesResponse.Count)

	pagewg := new(sync.WaitGroup)
	racewg := new(sync.WaitGroup)

	for i := 0; i < len(racesResponse.Races); i++ {
		detailJobs <- racesResponse.Races[0].Name
		racewg.Add(1)
	}

	pagewg.Add(racesResponse.NumPages - 1)

	for w := 0; w <= 5; w++ {
		go getPageWorker(w, jobs, detailJobs, pagewg, racewg)
	}

	for page := 2; page <= racesResponse.NumPages; page++ {
		jobs <- page
	}
	close(jobs)

	pagewg.Wait()

	for w := 0; w <= 5; w++ {
		go getRaceWorker(w, detailJobs, results, racewg)
	}

	racewg.Wait()
	//var sum = 0

	for raceDetail := range results {
		fmt.Println(raceDetail.Name, " ", raceDetail.Entrants[0].FinishTime)
		//sum += raceDetail.Entrants[0].FinishTime
	}
}
