package services

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
	"github.com/nadirakdag/finance-tracker/internal/storage"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type SummaryService struct {
	store  storage.Storage
	logger *logger.Logger
}

func NewSummaryService(store storage.Storage, logger *logger.Logger) *SummaryService {
	return &SummaryService{
		store:  store,
		logger: logger,
	}
}

func (s *SummaryService) GetSummary() (*models.Summary, error) {
	expenses, err := s.store.GetExpenses()
	if err != nil {
		return nil, err
	}

	incomes, err := s.store.GetIncomes()
	if err != nil {
		return nil, err
	}

	summary := &models.Summary{
		MonthlyBreakdown:  make(map[string]models.MonthData),
		CategoryBreakdown: make(map[string]models.CategoryData),
	}

	// Calculate totals and breakdowns
	for _, expense := range expenses {
		summary.TotalExpenses += expense.Amount

		monthKey := expense.Date.Format("2006-01")
		monthData := summary.MonthlyBreakdown[monthKey]
		monthData.Expenses += expense.Amount
		monthData.Balance -= expense.Amount
		summary.MonthlyBreakdown[monthKey] = monthData

		categoryData := summary.CategoryBreakdown[expense.Category]
		categoryData.Amount += expense.Amount
		categoryData.Type = "expense"
		summary.CategoryBreakdown[expense.Category] = categoryData
	}

	for _, income := range incomes {
		summary.TotalIncome += income.Amount
		monthKey := income.Date.Format("2006-01")
		monthData := summary.MonthlyBreakdown[monthKey]
		monthData.Income += income.Amount
		monthData.Balance += income.Amount
		summary.MonthlyBreakdown[monthKey] = monthData
	}

	summary.RemainingAmount = summary.TotalIncome - summary.TotalExpenses
	return summary, nil
}
