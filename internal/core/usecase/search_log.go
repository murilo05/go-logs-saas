package usecase

import (
	"context"
	"go-log-saas/internal/core/domain"
)

func (uc *ingestUseCase) Search(ctx context.Context, id string) (domain.IngestOutput, error) {
	return domain.IngestOutput{}, nil
}
