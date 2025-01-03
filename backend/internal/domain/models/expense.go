package models

import (
	"time"

	"github.com/nadirakdag/finance-tracker/internal/domain/errors"
)

type Expense struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (e *Expense) Validate() error {
	if e.Amount <= 0 {
		return errors.ErrInvalidAmount
	}
	if e.Description == "" {
		return errors.ErrEmptyDescription
	}
	if e.Date.IsZero() || e.Date.After(time.Now()) {
		return errors.ErrInvalidDate
	}
	if e.Category == "" {
		return errors.ErrInvalidCategory
	}
	return nil
}
