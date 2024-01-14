package repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/yramanovich/runestones/runestones"
)

func NewPostgres(ctx context.Context) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Postgres{db: pool}, nil
}

type Postgres struct {
	db *pgxpool.Pool
}

func (pg *Postgres) CreateRunestone(ctx context.Context, url string) (string, error) {
	row := pg.db.QueryRow(ctx, "INSERT INTO runestones(url) VALUES ($1) RETURNING id", url)
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (pg *Postgres) FindRunestone(ctx context.Context, id string) (runestones.Runestone, error) {
	row := pg.db.QueryRow(ctx, "SELECT r.id, r.url, r.created_time FROM runestones r WHERE r.id=$1", id)
	var runestone runestones.Runestone
	if err := row.Scan(&runestone.Id, &runestone.Url, &runestone.CreatedTime); err != nil {
		return runestone, err
	}
	return runestone, nil
}
