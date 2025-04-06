package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	dbpool  *pgxpool.Pool
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

		queryArgs := pgx.NamedArgs{
			"raceName": <-params.Jobs,
		}

		var race RaceRecord
		queryError := params.dbpool.QueryRow(context.Background(), `SELECT * FROM Races WHERE name = @raceName`, queryArgs).Scan(&race)

		if queryError == nil {
			fmt.Print("Record already exists. Skipping...")

			fmt.Println("Did not need to process race ", race.name)

			params.Wg.Done()
			return
		}

		response := GetRaceDetails(j)

		insertArgs := pgx.NamedArgs{
			"name":              response.Name,
			"categoryName":      response.Category.Name,
			"categoryShortName": response.Category.ShortName,
			"url":               response.Url,
			"goalName":          response.Goal.Name,
			"startedAt":         response.StartedAt,
		}

		_, insertError := params.dbpool.Exec(context.Background(), `INSERT INTO Races (name, category_name, category_short_name, url, goal_name, started_at) VALUES (@name, @categoryName, @categoryShortName, @url, @goalName, @startedAt)`, insertArgs)

		if insertError != nil {
			fmt.Print("Failed to insert record. Skipping...")
			params.Wg.Done()
			return
		}

		fmt.Println("Successfully processed new race ", response.Name)
		params.Results <- response
		params.Wg.Done()
	}
}
