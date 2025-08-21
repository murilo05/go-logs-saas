package usecase

import (
	"context"
	"go-log-saas/internal/core/domain"
)

type IngestUseCase interface {
	Ingest(ctx context.Context, ingest domain.Ingest) (domain.IngestOutput, error)
	Search(ctx context.Context, id string) (domain.IngestOutput, error)
}
