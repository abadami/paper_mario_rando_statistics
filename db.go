package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func CreatePool() (*pgxpool.Pool, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=randomizer_stats", os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PWD"), os.Getenv("POSTGRES_PORT"))

	dbpool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		return nil, err
	}

	pool = dbpool

	return dbpool, nil
}

func GetRaceByName(queryArgs pgx.NamedArgs) (string, error) {
	var race string
	queryError := pool.QueryRow(context.Background(), `SELECT name FROM Races WHERE name = @raceName`, queryArgs).Scan(&race)

	if queryError != nil {
		return "", queryError
	}

	return race, nil
}

func GetRacesByRaceEntrant(queryArgs pgx.NamedArgs) ([]RaceEntrantAndRaceRecord, error) {
	rows, _ := pool.Query(context.Background(), `SELECT 
	 RaceEntrants.id,
	 RaceEntrants.race_id,
	 RaceEntrants.entrant_id,
	 RaceEntrants.finish_time,
	 RaceEntrants.place,
	 RaceEntrants.place_ordinal,
	 RaceEntrants.status,
	 Races.name,
	 Races.category_name,
	 Races.category_short_name,
	 Races.url,
	 Races.goal_name,
	 Races.started_at FROM RaceEntrants 
	 	LEFT JOIN Races ON RaceEntrants.race_id = Races.id 
	 	WHERE RaceEntrants.entrant_id = @entrantId and goal_name = 'Blitz / 4 Chapters LCL Beat Bowser'`, queryArgs)

	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[RaceEntrantAndRaceRecord])

	if err != nil {
		return nil, err
	}

	return records, nil
}

func GetEntrant(queryArgs pgx.NamedArgs) (int, error) {
	var entrant_id int
	queryError := pool.QueryRow(context.Background(), `SELECT id FROM Entrants WHERE name = @entrantName`, queryArgs).Scan(&entrant_id)

	if queryError != nil {
		return -1, queryError
	}

	return entrant_id, nil
}

func GetEntrants() ([]EntrantRecord, error) {
	rows, _ := pool.Query(context.Background(), `SELECT * FROM Entrants`)

	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[EntrantRecord])

	if err != nil {
		return nil, err
	}

	return records, nil
}

func InsertRaceDetails(details RaceDetail) error {
	insertArgs := pgx.NamedArgs{
		"name":              details.Name,
		"categoryName":      details.Category.Name,
		"categoryShortName": details.Category.ShortName,
		"url":               details.Url,
		"goalName":          details.Goal.Name,
		"startedAt":         details.StartedAt,
	}

	var raceId int
	insertRaceError := pool.QueryRow(context.Background(), `INSERT INTO Races (name, category_name, category_short_name, url, goal_name, started_at) VALUES (@name, @categoryName, @categoryShortName, @url, @goalName, @startedAt) RETURNING id`, insertArgs).Scan(&raceId)

	if insertRaceError != nil {
		return insertRaceError
	}

	for _, entrant := range details.Entrants {
		insertEntrantArgs := pgx.NamedArgs{
			"name": entrant.User.Name,
			"fullName": entrant.User.FullName,
		}


		var entrantId int

		//Check if entrant already exists
		findEntrantError := pool.QueryRow(context.Background(), `SELECT id FROM Entrants WHERE full_name=@fullName`, insertEntrantArgs).Scan(&entrantId)

		//If entrant doesn't exist, insert entrant. If weird error happens, continue I guess? Need better handling here still (I think fail entire task and add to tasklog would be correct)
		if findEntrantError == pgx.ErrNoRows {
			insertEntrantError := pool.QueryRow(context.Background(), `INSERT INTO Entrants (name, full_name) VALUES (@name, @fullName) RETURNING id`, insertEntrantArgs).Scan(&entrantId)

			if insertEntrantError != nil {
				fmt.Println("Failure to insert entrant: ", entrant.User.Name, " because ", insertEntrantError)
				continue
			}
		} else if findEntrantError != nil {
			fmt.Println("Weird error finding entrant: ", findEntrantError)
			continue
		}

		insertRaceEntrantArgs := pgx.NamedArgs{
			"raceId": raceId,
			"entrantId": entrantId,
			"finishTime": entrant.FinishTime,
			"place": entrant.Place,
			"placeOrdinal": entrant.PlaceOrdinal,
			"status": entrant.Status.Value,
		}

		_, insertRaceEntrantError := pool.Exec(context.Background(), `INSERT INTO RaceEntrants (race_id, entrant_id, finish_time, place, place_ordinal, status) VALUES (@raceId, @entrantId, @finishTime, @place, @placeOrdinal, @status)`, insertRaceEntrantArgs)

		if insertRaceEntrantError != nil {
			fmt.Println("Failure to insert entrant race details because ", insertRaceEntrantError)
			continue
		}
	}
	

	return nil
}

func InsertTaskLog(success bool, racesFetched int) error {
	insertTaskLogArgs := pgx.NamedArgs{
		"dateRan":      time.Now().Format(time.RFC3339),
		"racesFetched": racesFetched,
		"successful":   success,
	}

	_, taskLogError := pool.Exec(context.Background(), `INSERT INTO TaskLog (date_ran, races_fetched, successful) VALUES (@dateRan, @racesFetched, @successful)`, insertTaskLogArgs)

	return taskLogError
}

