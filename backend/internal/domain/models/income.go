package models

import (
	"time"
)

type Income struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Source      string    `json:"source"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
