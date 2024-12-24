package handlers

import (
	"encoding/json"
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

func (h *SummaryHandler) Get(w http.ResponseWriter, r *http.Request) {
	summary, err := h.service.GetSummary()
	if err != nil {
		h.logger.Error("Failed to get summary", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(summary)
}
