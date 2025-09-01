package repository

import (
	"context"
	"go-log-saas/internal/core/domain"

	pg "go-log-saas/internal/repository/client"

	"go.uber.org/zap"
)

type repository struct {
	logger *zap.SugaredLogger
	pg     pg.Postgres
}

func NewRepository(pg pg.Postgres, logger *zap.SugaredLogger) IngestRepository {
	return &repository{
		logger: logger,
		pg:     pg,
	}
}

func (r repository) Save(ctx context.Context, ingest domain.Ingest) (domain.IngestOutput, error) {
	//TODO
	_, _ = r.pg.Save(ctx, ingest)

	return domain.IngestOutput{}, nil
}

func (r repository) Get(ctx context.Context, id string) (domain.IngestOutput, error) {
	//TODO
	return domain.IngestOutput{}, nil
}
