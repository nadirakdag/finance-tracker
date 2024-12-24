package validation

import (
	"errors"
	"github.com/nadirakdag/finance-tracker/internal/domain/models"
)

var (
	ErrInvalidAmount      = errors.New("amount must be greater than 0")
	ErrMissingCategory    = errors.New("category is required")
	ErrMissingDescription = errors.New("description is required")
	ErrInvalidDate        = errors.New("invalid date")
)

func ValidateExpense(expense *models.Expense) error {
	if expense.Amount <= 0 {
		return ErrInvalidAmount
	}
	if expense.Category == "" {
		return ErrMissingCategory
	}
	if expense.Description == "" {
		return ErrMissingDescription
	}
	if expense.Date.IsZero() {
		return ErrInvalidDate
	}
	return nil
}

func ValidateIncome(income *models.Income) error {
	if income.Amount <= 0 {
		return ErrInvalidAmount
	}
	if income.Source == "" {
		return ErrMissingCategory
	}
	if income.Description == "" {
		return ErrMissingDescription
	}
	if income.Date.IsZero() {
		return ErrInvalidDate
	}
	return nil
}
