package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool() (*pgxpool.Pool, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=randomizer_stats", os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PWD"), os.Getenv("POSTGRES_PORT"))

	dbpool, err := pgxpool.New(context.Background(), url)

	if err != nil {
		return nil, err
	}

	return dbpool, nil
}