package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/nadirakdag/finance-tracker/pkg/logger"
)

type HealthHandler struct {
	logger    *logger.Logger
	startTime time.Time
}

type HealthResponse struct {
	Status    string   `json:"status"`
	Uptime    string   `json:"uptime"`
	GoVersion string   `json:"goVersion"`
	Memory    MemStats `json:"memory"`
}

type MemStats struct {
	Alloc      uint64 `json:"alloc"`      // bytes allocated and still in use
	TotalAlloc uint64 `json:"totalAlloc"` // bytes allocated (even if freed)
	Sys        uint64 `json:"sys"`        // bytes obtained from system
	NumGC      uint32 `json:"numGC"`      // number of completed GC cycles
}

func NewHealthHandler(logger *logger.Logger) *HealthHandler {
	return &HealthHandler{
		logger:    logger,
		startTime: time.Now(),
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	health := HealthResponse{
		Status:    "ok",
		Uptime:    time.Since(h.startTime).String(),
		GoVersion: runtime.Version(),
		Memory: MemStats{
			Alloc:      memStats.Alloc,
			TotalAlloc: memStats.TotalAlloc,
			Sys:        memStats.Sys,
			NumGC:      memStats.NumGC,
		},
	}

	writeJSON(w, http.StatusOK, health)
}
