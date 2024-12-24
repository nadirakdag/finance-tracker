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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateIncome(&income); err != nil {
		h.logger.Error("Failed to create income", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(income)
}

func (h *IncomeHandler) List(w http.ResponseWriter, r *http.Request) {
	incomes, err := h.service.GetIncomes()
	if err != nil {
		h.logger.Error("Failed to get incomes", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(incomes)
}
