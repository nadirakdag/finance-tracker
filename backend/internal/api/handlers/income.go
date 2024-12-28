package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nadirakdag/finance-tracker/internal/domain/errors"
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type IncomeHandler struct {
	service         *services.IncomeService
	categoryService *services.CategoryService
	logger          *logger.Logger
}

func NewIncomeHandler(service *services.IncomeService, categoryService *services.CategoryService, logger *logger.Logger) *IncomeHandler {
	return &IncomeHandler{
		service:         service,
		categoryService: categoryService,
		logger:          logger,
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

	//validate expense category
	if err := h.categoryService.CheckCategory(income.Source, "income"); err != nil {
		h.logger.Warn("Invalid expense data", "error", err)
		writeError(w, errors.ErrInvalidCategory)
		return
	}

	if err := h.service.CreateIncome(&income); err != nil {
		h.logger.Error("Failed to create income", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, income)
}

func (h *IncomeHandler) List(w http.ResponseWriter, _ *http.Request) {
	incomes, err := h.service.GetIncomes()
	if err != nil {
		h.logger.Error("Failed to get incomes", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, incomes)
}
