package main

import (
	"fmt"
	"sync"
)

type GetPageWorkerParams struct {
	Id int
	Jobs <-chan int
	Results chan<- string
	Pagewg *sync.WaitGroup
	Racewg *sync.WaitGroup
}

type GetRaceWorkerParams struct {
	Id int
	Jobs <-chan string
	Results chan<- RaceDetail
	Wg *sync.WaitGroup
}

func GetPageWorker(params GetPageWorkerParams) {
	for j := range params.Jobs {
		fmt.Println("getPageWorker", params.Id, "started job", j)
		response := GetRaceTitlesAndEntrantsByPage(j)
		for race := range response.Races {
			fmt.Println("Processing race title ", response.Races[race].Name)
			params.Racewg.Add(1)
			params.Results <- response.Races[race].Name
		}
		params.Pagewg.Done()
	}
}

func GetRaceWorker(params GetRaceWorkerParams) {
	for j := range params.Jobs {
		fmt.Println("getRaceWorker", params.Id, "started job", j)
		response := GetRaceDetails(j)
		fmt.Println("Processed race ", response.Name)
		params.Results <- response
		params.Wg.Done()
	}
}