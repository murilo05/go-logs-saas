package usecase

import (
	"context"
	"errors"
	"go-log-saas/internal/core/domain"
	"go-log-saas/internal/core/utils"
	"go-log-saas/internal/repository"

	"go.uber.org/zap"
)

type ingestUseCase struct {
	logger *zap.SugaredLogger
	repo   repository.IngestRepository
}

func NewIngestUseCase(logger *zap.SugaredLogger, repo repository.IngestRepository) IngestUseCase {
	return &ingestUseCase{
		logger: logger,
	}
}

func (uc *ingestUseCase) Ingest(ctx context.Context, ingestion domain.Ingest) (domain.IngestOutput, error) {
	uc.logger.Warnw("UseCase: Ingest Started", "ingest_id", ingestion.ID)

	if !utils.IsValidLevel(ingestion.Level) {
		err := errors.New("invalid log level")
		uc.logger.Error("failed to validate log level: ", err)
		return domain.IngestOutput{}, err
	}

	output, err := uc.repo.Save(ctx, ingestion)
	if err != nil {
		uc.logger.Error("repository failed to save ingestion: ", err)
		return domain.IngestOutput{}, err
	}

	return output, nil
}
