package postgres

import (
	"context"
	"go-log-saas/internal/core/domain"
	pg "go-log-saas/internal/repository/client"
	"log"
)

var _ pg.Postgres = &PG{}

func (c PG) Save(ctx context.Context, ingest domain.Ingest) (domain.IngestOutput, error) {
	//TODO INSERT

	_, err := c.Exec(context.Background(), "")
	if err != nil {
		log.Fatal("Error while inserting value into the table")
	}
	return domain.IngestOutput{}, nil
}

func (c PG) Get(ctx context.Context, id string) (domain.IngestOutput, error) {
	//TODO SELECT
	return domain.IngestOutput{}, nil
}
