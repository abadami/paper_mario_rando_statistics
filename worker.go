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
		fmt.Println("Page: ", j)
		response := GetRaceTitlesAndEntrantsByPage(j)
		for race := range response.Races {
			params.Racewg.Add(1)
			params.Results <- response.Races[race].Name
		}
		params.Pagewg.Done()
	}
}

func GetRaceWorker(params GetRaceWorkerParams) {
	for j := range params.Jobs {
		func () {
			defer params.Wg.Done()

			queryArgs := pgx.NamedArgs{
				"raceName": j,
			}
	
			var race string
			queryError := params.dbpool.QueryRow(context.Background(), `SELECT name FROM Races WHERE name = @raceName`, queryArgs).Scan(&race)
	
			//TODO: More specific error handling
			if queryError == nil {
				fmt.Println("Found Race ", race)
				return
			}

			if queryError != pgx.ErrNoRows {
				fmt.Println("Weird Error", queryError)
			}

			fmt.Println("Did not find race, ", j)
	
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
				return
			}
	
			params.Results <- response
		}()
	}
}
