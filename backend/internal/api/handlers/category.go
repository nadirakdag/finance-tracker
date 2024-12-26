package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type CategoryHandler struct {
	service *services.CategoryService
	logger  *logger.Logger
}

func NewCategoryHandler(service *services.CategoryService, logger *logger.Logger) *CategoryHandler {
	return &CategoryHandler{
		service: service,
		logger:  logger,
	}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryType := vars["type"]

	categories, err := h.service.GetCategories(categoryType)
	if err != nil {
		h.logger.Error("Failed to get categories", "error", err)
		http.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
