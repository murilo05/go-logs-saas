package usecase

import (
	"context"
	"go-log-saas/internal/core/domain"
)

type IngestUseCase interface {
	Ingest(context.Context, domain.Ingest) (domain.IngestOutput, error)
}
