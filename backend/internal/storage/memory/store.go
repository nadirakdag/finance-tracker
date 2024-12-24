package memory

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/pkg/utils"
	"sync"
)

type memoryStore struct {
	expenses []models.Expense
	incomes  []models.Income
	mutex    sync.RWMutex
}

func NewStore() *memoryStore {
	return &memoryStore{
		expenses: make([]models.Expense, 0),
		incomes:  make([]models.Income, 0),
	}
}

func (s *memoryStore) CreateExpense(expense *models.Expense) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Generate ID if not provided
	if expense.ID == "" {
		expense.ID = utils.GenerateID(16) // 16 characters long
	}

	s.expenses = append(s.expenses, *expense)
	return nil
}

func (s *memoryStore) GetExpenses() ([]models.Expense, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	expenses := make([]models.Expense, len(s.expenses))
	copy(expenses, s.expenses)
	return expenses, nil
}

func (s *memoryStore) CreateIncome(income *models.Income) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Generate ID if not provided
	if income.ID == "" {
		income.ID = utils.GenerateID(16) // 16 characters long
	}

	s.incomes = append(s.incomes, *income)
	return nil
}

func (s *memoryStore) GetIncomes() ([]models.Income, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	incomes := make([]models.Income, len(s.incomes))
	copy(incomes, s.incomes)
	return incomes, nil
}
