package errors

import "errors"

var (
	ErrInvalidAmount    = errors.New("amount must be greater than 0")
	ErrEmptyDescription = errors.New("description cannot be empty")
	ErrInvalidDate      = errors.New("invalid date")
	ErrInvalidCategory  = errors.New("invalid category")
	ErrNotFound         = errors.New("resource not found")
)

type ErrorResponse struct {
	Error       string `json:"error"`
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
}
