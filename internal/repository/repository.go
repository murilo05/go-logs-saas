package repository

import (
	"context"
	"go-log-saas/internal/core/domain"

	"go.uber.org/zap"
)

type repository struct {
	logger *zap.SugaredLogger
}

func NewRepository(logger *zap.SugaredLogger) IngestRepository {
	return &repository{
		logger: logger,
	}
}

func (r repository) Save(ctx context.Context, ingest domain.Ingest) (domain.IngestOutput, error) {
	return domain.IngestOutput{}, nil
}

func (r repository) Get(ctx context.Context, id string) (domain.IngestOutput, error) {
	return domain.IngestOutput{}, nil
}
