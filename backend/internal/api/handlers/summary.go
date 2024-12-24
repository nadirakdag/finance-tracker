package handlers

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
	"net/http"
)

type SummaryHandler struct {
	service *services.SummaryService
	logger  *logger.Logger
}

func NewSummaryHandler(service *services.SummaryService, logger *logger.Logger) *SummaryHandler {
	return &SummaryHandler{
		service: service,
		logger:  logger,
	}
}

func (h *SummaryHandler) Get(w http.ResponseWriter) {
	summary, err := h.service.GetSummary()
	if err != nil {
		h.logger.Error("Failed to get summary", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, summary)
}
