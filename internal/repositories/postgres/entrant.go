package postgres

import (
	"context"

	"github.com/abadami/randomizer-statistics/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EntrantRepository struct {
	pool *pgxpool.Pool
}

func NewEntrantRepository(pool *pgxpool.Pool) *EntrantRepository {
	return &EntrantRepository{pool}
}

func (repo *EntrantRepository) GetEntrant(queryArgs pgx.NamedArgs) (int, error) {
	var entrant_id int
	queryError := repo.pool.QueryRow(context.Background(), `SELECT id FROM Entrants WHERE name = @entrantName`, queryArgs).Scan(&entrant_id)

	if queryError != nil {
		return -1, queryError
	}

	return entrant_id, nil
}

func (repo *EntrantRepository) GetEntrants() ([]domain.EntrantRecord, error) {
	rows, _ := repo.pool.Query(context.Background(), `SELECT * FROM Entrants`)

	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.EntrantRecord])

	if err != nil {
		return nil, err
	}

	return records, nil
}