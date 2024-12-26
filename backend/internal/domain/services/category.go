package services

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/storage"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type CategoryService struct {
	store  storage.Storage
	logger *logger.Logger
}

func NewCategoryService(store storage.Storage, logger *logger.Logger) *CategoryService {
	return &CategoryService{
		store:  store,
		logger: logger,
	}
}
func (s *CategoryService) GetCategories(cType string) ([]models.Category, error) {
	return s.store.GetCategories(cType)
}
