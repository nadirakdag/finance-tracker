package services

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/storage"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type ExpenseService struct {
	store  storage.Storage
	logger *logger.Logger
}

func NewExpenseService(store storage.Storage, logger *logger.Logger) *ExpenseService {
	return &ExpenseService{
		store:  store,
		logger: logger,
	}
}

func (s *ExpenseService) CreateExpense(expense *models.Expense) error {
	return s.store.CreateExpense(expense)
}

func (s *ExpenseService) GetExpenses() ([]models.Expense, error) {
	return s.store.GetExpenses()
}
