package main

import (
	"fmt"
	"sync"
)

func getPageWorker(id int, jobs <-chan int, results chan<- string, pagewg *sync.WaitGroup, racewg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("getPageWorker", id, "started job", j)
		response := GetRaceTitlesAndEntrantsByPage(j)
		for race := range response.Races {
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
		response := GetRaceDetails(j)
		fmt.Println("Processed race ", response.Name)
		results <- response
		wg.Done()
	}
}

func main() {
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
		go getPageWorker(w, jobs, detailJobs, pagewg, racewg)
	}

	for page := 2; page <= racesResponse.NumPages; page++ {
		jobs <- page
	}
	close(jobs)

	for w := 0; w <= 100; w++ {
		go getRaceWorker(w, detailJobs, results, racewg)
	}

	pagewg.Wait()

	//We know we have all of the detailJobs, so we can close here
	close(detailJobs)

	racewg.Wait()

	//We know we have all the results now, so close the channels
	close(results)

	sum := 0
	count := 0

	for raceDetail := range results {
		fmt.Println(raceDetail.Name, " ", raceDetail.Entrants[0].FinishTime)
		sum += ParseTimeString(raceDetail.Entrants[0].FinishTime)
		count += 1
	}

	fmt.Println(count)
	fmt.Println(ParseSecondsToTime(sum / count))
}
