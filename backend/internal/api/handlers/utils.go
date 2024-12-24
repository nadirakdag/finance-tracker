package handlers

import (
	"encoding/json"
	"github.com/nadirakdag/finance-tracker/internal/domain/errors"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	var status int
	var response errors.ErrorResponse

	switch err {
	case errors.ErrInvalidAmount, errors.ErrEmptyDescription, errors.ErrInvalidDate, errors.ErrInvalidCategory:
		status = http.StatusBadRequest
		response = errors.ErrorResponse{
			Error:       "ValidationError",
			Code:        "INVALID_INPUT",
			Description: err.Error(),
		}
	case errors.ErrNotFound:
		status = http.StatusNotFound
		response = errors.ErrorResponse{
			Error:       "NotFoundError",
			Code:        "RESOURCE_NOT_FOUND",
			Description: err.Error(),
		}
	default:
		status = http.StatusInternalServerError
		response = errors.ErrorResponse{
			Error:       "InternalError",
			Code:        "INTERNAL_ERROR",
			Description: "An internal error occurred",
		}
	}

	writeJSON(w, status, response)
}
