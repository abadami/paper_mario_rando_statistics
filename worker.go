package main

import (
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

type GetPageWorkerParams struct {
	Id      int
	Jobs    <-chan int
	Results chan<- string
	Pagewg  *sync.WaitGroup
	Racewg  *sync.WaitGroup
}

type GetRaceWorkerParams struct {
	Id      int
	Jobs    <-chan string
	Results chan<- RaceDetail
	Wg      *sync.WaitGroup
}

func GetPageWorker(params GetPageWorkerParams) {
	for j := range params.Jobs {
		response := GetRaceTitlesAndEntrantsByPage(j)
		for race := range response.Races {
			params.Racewg.Add(1)
			params.Results <- response.Races[race].Name
		}
		params.Pagewg.Done()
	}
}

func GetRaceWorker(params GetRaceWorkerParams) {
	for job := range params.Jobs {
			queryArgs := pgx.NamedArgs{
				"raceName": job,
			}
	
			_, queryError := GetRaceByName(queryArgs)
	
			//TODO: More specific error handling
			if queryError == nil {
				params.Wg.Done()
				continue
			}

			if queryError != pgx.ErrNoRows {
				fmt.Println("Weird Error", queryError)
				params.Wg.Done()
				continue
			}
	
			response := GetRaceDetails(job)
	
			insertRaceError := InsertRaceDetails(response)
	
			if insertRaceError != nil {
				params.Wg.Done()
				continue
			}
	
			params.Results <- response
			params.Wg.Done()
	}
}
