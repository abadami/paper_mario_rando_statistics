package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/abadami/randomizer-statistics/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RaceRepository struct {
	pool *pgxpool.Pool
}

func NewRaceRepository(pool *pgxpool.Pool) *RaceRepository {
	return &RaceRepository{pool}
}

func (repo *RaceRepository) GetRaceByName(queryArgs pgx.NamedArgs) (string, error) {
	var race string
	queryError := repo.pool.QueryRow(context.Background(), `SELECT name FROM Races WHERE name = @raceName`, queryArgs).Scan(&race)

	if queryError != nil {
		return "", queryError
	}

	return race, nil
}

func (repo *RaceRepository) GetRacesByRaceEntrant(request domain.StatisticsRequest) ([]domain.RaceEntrantAndRaceRecord, error) {
	var queryBuilder strings.Builder
	
	var queryArgs = pgx.NamedArgs{
		"goal": request.Goal,
		"entrantId": request.ContainsEntrant,
	}

	queryBuilder.WriteString(`SELECT 
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
		WHERE goal_name = @goal`)

	fmt.Printf("%d", request.ContainsEntrant)

	if (request.ContainsEntrant > -1) {
		queryBuilder.WriteString(" and Races.id in (Select RaceEntrants.race_id From RaceEntrants where RaceEntrants.entrant_id = @entrantId)")
	}

	if (request.RaceType == "league") {
		queryBuilder.WriteString(" and (SELECT COUNT(*) FROM RaceEntrants WHERE race_id = Races.id) = 2")
	}

	if (request.RaceType == "community") {
		queryBuilder.WriteString(" and (SELECT COUNT(*) FROM RaceEntrants WHERE race_id = Races.id) > 2")
	}

	rows, _ := repo.pool.Query(context.Background(), queryBuilder.String(), queryArgs)

	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.RaceEntrantAndRaceRecord])

	if err != nil {
		return nil, err
	}

	return records, nil
}

//TODO: Move entrant insertion to entrant.go
func (repo *RaceRepository)  InsertRaceDetails(details domain.RaceDetail) error {
	insertArgs := pgx.NamedArgs{
		"name":              details.Name,
		"categoryName":      details.Category.Name,
		"categoryShortName": details.Category.ShortName,
		"url":               details.Url,
		"goalName":          details.Goal.Name,
		"startedAt":         details.StartedAt,
	}

	var raceId int
	insertRaceError := repo.pool.QueryRow(context.Background(), `INSERT INTO Races (name, category_name, category_short_name, url, goal_name, started_at) VALUES (@name, @categoryName, @categoryShortName, @url, @goalName, @startedAt) RETURNING id`, insertArgs).Scan(&raceId)

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
		findEntrantError := repo.pool.QueryRow(context.Background(), `SELECT id FROM Entrants WHERE full_name=@fullName`, insertEntrantArgs).Scan(&entrantId)

		//If entrant doesn't exist, insert entrant. If weird error happens, continue I guess? Need better handling here still (I think fail entire task and add to tasklog would be correct)
		if findEntrantError == pgx.ErrNoRows {
			insertEntrantError := repo.pool.QueryRow(context.Background(), `INSERT INTO Entrants (name, full_name) VALUES (@name, @fullName) RETURNING id`, insertEntrantArgs).Scan(&entrantId)

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

		_, insertRaceEntrantError := repo.pool.Exec(context.Background(), `INSERT INTO RaceEntrants (race_id, entrant_id, finish_time, place, place_ordinal, status) VALUES (@raceId, @entrantId, @finishTime, @place, @placeOrdinal, @status)`, insertRaceEntrantArgs)

		if insertRaceEntrantError != nil {
			fmt.Println("Failure to insert entrant race details because ", insertRaceEntrantError)
			continue
		}
	}
	

	return nil
}