package handlers

import (
	"encoding/json"
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
	"net/http"
)

type ExpenseHandler struct {
	service *services.ExpenseService
	logger  *logger.Logger
}

func NewExpenseHandler(service *services.ExpenseService, logger *logger.Logger) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
		logger:  logger,
	}
}

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		h.logger.Error("Failed to decode expense", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateExpense(&expense); err != nil {
		h.logger.Error("Failed to create expense", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) List(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.service.GetExpenses()
	if err != nil {
		h.logger.Error("Failed to get expenses", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}
