package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nadirakdag/finance-tracker/internal/config"

	"github.com/nadirakdag/finance-tracker/internal/api"
	"github.com/nadirakdag/finance-tracker/internal/storage/sqlite"
	"github.com/nadirakdag/finance-tracker/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize appLog
	appLog := logger.NewLogger(&cfg.Logging)

	// Initialize storage
	store, err := sqlite.NewStore("./finance.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer store.Close()

	// Initialize router with all dependencies
	router := api.NewRouter(cfg, store, appLog)

	// Add prometheus metrics endpoint if enabled
	if cfg.Metrics.Enabled {
		router.Handle(cfg.Metrics.Path, promhttp.Handler())
	}

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	appLog.Info("Starting server", "address", serverAddr)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		appLog.Fatal("Server failed to start", "error", err)
	}
}
