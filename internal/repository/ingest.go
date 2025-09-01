package repository

import (
	"context"
	"go-log-saas/internal/core/domain"
)

type IngestRepository interface {
	Save(ctx context.Context, ingest domain.Ingest) (domain.IngestOutput, error)
	Get(ctx context.Context, id string) (domain.IngestOutput, error)
}
