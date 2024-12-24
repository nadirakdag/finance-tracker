package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadirakdag/finance-tracker/internal/api/handlers"
	"github.com/nadirakdag/finance-tracker/internal/api/middleware"
	"github.com/nadirakdag/finance-tracker/internal/domain/services"
	"github.com/nadirakdag/finance-tracker/internal/storage"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type Router struct {
	router *mux.Router
	logger *logger.Logger
	store  storage.Storage
}

func NewRouter(store storage.Storage, logger *logger.Logger) *mux.Router {
	router := mux.NewRouter()

	// Initialize services
	expenseService := services.NewExpenseService(store, logger)
	incomeService := services.NewIncomeService(store, logger)
	summaryService := services.NewSummaryService(store, logger)

	// Initialize handlers
	expenseHandler := handlers.NewExpenseHandler(expenseService, logger)
	incomeHandler := handlers.NewIncomeHandler(incomeService, logger)
	summaryHandler := handlers.NewSummaryHandler(summaryService, logger)

	// API Routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Expenses endpoints
	api.HandleFunc("/expenses", expenseHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/expenses", expenseHandler.List).Methods(http.MethodGet)

	// Incomes endpoints
	api.HandleFunc("/incomes", incomeHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/incomes", incomeHandler.List).Methods(http.MethodGet)

	// Summary endpoint
	api.HandleFunc("/summary", summaryHandler.Get).Methods(http.MethodGet)

	// Add middleware
	router.Use(middleware.Cors)
	router.Use(middleware.Logging(logger))

	return router
}
