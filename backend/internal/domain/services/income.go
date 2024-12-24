package services

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/storage"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type IncomeService struct {
	store  storage.Storage
	logger *logger.Logger
}

func NewIncomeService(store storage.Storage, logger *logger.Logger) *IncomeService {
	return &IncomeService{
		store:  store,
		logger: logger,
	}
}

func (s *IncomeService) CreateIncome(income *models.Income) error {
	return s.store.CreateIncome(income)
}

func (s *IncomeService) GetIncomes() ([]models.Income, error) {
	return s.store.GetIncomes()
}
