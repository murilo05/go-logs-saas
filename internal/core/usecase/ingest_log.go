package usecase

import (
	"context"
	"go-log-saas/internal/core/domain"

	"go.uber.org/zap"
)

type ingestUseCase struct {
	logger *zap.SugaredLogger
}

func NewIngestUseCase(logger *zap.SugaredLogger) IngestUseCase {
	return &ingestUseCase{
		logger: logger,
	}
}

func (uc *ingestUseCase) Ingest(ctx context.Context, ingestion domain.Ingest) (domain.IngestOutput, error) {
	return domain.IngestOutput{}, nil
}
