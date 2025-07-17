package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskLogRepository struct {
	pool *pgxpool.Pool
}

func NewTaskLogRepository(pool *pgxpool.Pool) *TaskLogRepository {
	return &TaskLogRepository{pool}
}

func (repo *TaskLogRepository) InsertTaskLog(success bool, racesFetched int) error {
	insertTaskLogArgs := pgx.NamedArgs{
		"dateRan":      time.Now().Format(time.RFC3339),
		"racesFetched": racesFetched,
		"successful":   success,
	}

	_, taskLogError := repo.pool.Exec(context.Background(), `INSERT INTO TaskLog (date_ran, races_fetched, successful) VALUES (@dateRan, @racesFetched, @successful)`, insertTaskLogArgs)

	return taskLogError
}