package models

import (
	"github.com/nadirakdag/finance-tracker/internal/domain/errors"
	"time"
)

var validIncomeSources = map[string]bool{
	"salary":    true,
	"freelance": true,
	"dividend":  true,
	"gift":      true,
	"other":     true,
}

type Income struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Source      string    `json:"source"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (i *Income) Validate() error {
	if i.Amount <= 0 {
		return errors.ErrInvalidAmount
	}
	if i.Description == "" {
		return errors.ErrEmptyDescription
	}
	if i.Date.IsZero() || i.Date.After(time.Now()) {
		return errors.ErrInvalidDate
	}
	if !validIncomeSources[i.Source] {
		return errors.ErrInvalidCategory
	}
	return nil
}
