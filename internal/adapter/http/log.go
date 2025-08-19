package http

import (
	"go-log-saas/internal/core/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *usecase.LogService
}

func NewHandler(svc *usecase.LogService) *Handler {
	return &Handler{
		svc,
	}
}

func (h *Handler) IngestLog(gin *gin.Context) {

}

func (h *Handler) SearchLog(gin *gin.Context) {

}

func (h *Handler) SearchLogById(gin *gin.Context) {

}
