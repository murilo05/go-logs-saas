package http

import (
	"errors"
	"go-log-saas/internal/adapter/http/dto"
	"go-log-saas/internal/adapter/http/response"
	"go-log-saas/internal/core/domain"
	"go-log-saas/internal/core/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	uc     usecase.IngestUseCase
	logger *zap.SugaredLogger
}

func NewHandler(uc usecase.IngestUseCase, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		uc,
		logger,
	}
}

func (h *Handler) IngestLog(ctx *gin.Context) {
	var req dto.IngestInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error on bind JSON")
		response.HandleValidationError(ctx, err)
		return
	}

	if err := checkFields(req); err != nil {
		h.logger.Error("Error on field check")
		response.HandleValidationError(ctx, err)
		return
	}

	ingestion := domain.Ingest{
		ID:      uuid.New().String(),
		APIKey:  req.APIKey,
		AppID:   req.AppID,
		Level:   req.Level,
		Message: req.Message,
		Context: req.Context,
	}

	rsp, err := h.uc.Ingest(ctx, ingestion)
	if err != nil {
		h.logger.Error("UseCase Failed during ingestion")
		response.HandleError(ctx, err)
		return
	}

	h.logger.Infow("Log Ingested Successfully", "id", rsp.ID, "status", rsp.Status)
	response.HandleSuccess(ctx, rsp.ID, rsp.Status, nil)
}

func (h *Handler) SearchLogById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.logger.Error("Error on field check")
		response.HandleValidationError(ctx, errors.New("invalid log id"))
		return
	}

	rsp, err := h.uc.Search(ctx, id)
	if err != nil {
		h.logger.Error("UseCase Failed during ingestion")
		response.HandleError(ctx, err)
		return
	}

	h.logger.Infow("Log Ingested Successfully", "id", rsp.ID, "status", rsp.Status)
	response.HandleSuccess(ctx, rsp.ID, rsp.Status, &rsp)
}

func checkFields(req dto.IngestInput) error {
	if req.APIKey == "" {
		return errors.New("field 'api_key' can't be nil")
	}
	if req.AppID == "" {
		return errors.New("field 'app_id' can't be nil")
	}
	if req.Level == "" {
		return errors.New("field 'level' can't be nil")
	}
	if req.Message == "" {
		return errors.New("field 'message' can't be nil")
	}
	return nil
}
