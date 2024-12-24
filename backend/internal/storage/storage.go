package storage

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
)

type Storage interface {
	CreateExpense(*models.Expense) error
	GetExpenses() ([]models.Expense, error)
	CreateIncome(*models.Income) error
	GetIncomes() ([]models.Income, error)
}
