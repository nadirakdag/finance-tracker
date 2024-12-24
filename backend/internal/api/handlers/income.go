package handlers

import (
	"encoding/json"
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
	"net/http"
)

type IncomeHandler struct {
	service *services.IncomeService
	logger  *logger.Logger
}

func NewIncomeHandler(service *services.IncomeService, logger *logger.Logger) *IncomeHandler {
	return &IncomeHandler{
		service: service,
		logger:  logger,
	}
}

func (h *IncomeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var income models.Income
	if err := json.NewDecoder(r.Body).Decode(&income); err != nil {
		h.logger.Error("Failed to decode income", "error", err)
		writeError(w, err)
		return
	}

	// Validate income
	if err := income.Validate(); err != nil {
		h.logger.Warn("Invalid income data", "error", err)
		writeError(w, err)
		return
	}

	if err := h.service.CreateIncome(&income); err != nil {
		h.logger.Error("Failed to create income", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, income)
}

func (h *IncomeHandler) List(w http.ResponseWriter) {
	incomes, err := h.service.GetIncomes()
	if err != nil {
		h.logger.Error("Failed to get incomes", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, incomes)
}
