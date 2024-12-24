package main

import (
	"fmt"
	"github.com/nadirakdag/finance-tracker/internal/config"
	"log"
	"net/http"

	"github.com/nadirakdag/finance-tracker/internal/api"
	"github.com/nadirakdag/finance-tracker/internal/storage/memory"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize logger
	logger := logger.NewLogger(&cfg.Logging)

	// Initialize storage
	store := memory.NewStore()

	// Initialize router with all dependencies
	router := api.NewRouter(cfg, store, logger)

	// Add prometheus metrics endpoint if enabled
	if cfg.Metrics.Enabled {
		router.Handle(cfg.Metrics.Path, promhttp.Handler())
	}

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Info("Starting server", "address", serverAddr)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		logger.Fatal("Server failed to start", "error", err)
	}
}
