package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nadirakdag/finance-tracker/internal/domain/errors"
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type ExpenseHandler struct {
	service         *services.ExpenseService
	categoryService *services.CategoryService
	logger          *logger.Logger
}

func NewExpenseHandler(service *services.ExpenseService, categoryService *services.CategoryService, logger *logger.Logger) *ExpenseHandler {
	return &ExpenseHandler{
		service:         service,
		categoryService: categoryService,
		logger:          logger,
	}
}

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		h.logger.Error("Failed to decode expense", "error", err)
		writeError(w, err)
		return
	}

	// Validate expense
	if err := expense.Validate(); err != nil {
		h.logger.Warn("Invalid expense data", "error", err)
		writeError(w, err)
		return
	}

	//validate expense category
	if err := h.categoryService.CheckCategory(expense.Category, "expense"); err != nil {
		h.logger.Warn("Invalid expense data", "error", err)
		writeError(w, errors.ErrInvalidCategory)
		return
	}

	if err := h.service.CreateExpense(&expense); err != nil {
		h.logger.Error("Failed to create expense", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, expense)
}

func (h *ExpenseHandler) List(w http.ResponseWriter, _ *http.Request) {
	expenses, err := h.service.GetExpenses()
	if err != nil {
		h.logger.Error("Failed to get expenses", "error", err)
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, expenses)
}
