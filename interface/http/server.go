package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// HealthStatus represents the health of the application.
type HealthStatus struct {
	Status      string `json:"status"`
	Connected   bool   `json:"connected"`
	ServiceName string `json:"service_name"`
}

// StartServer starts a simple HTTP server for health checks and metrics.
// It runs in a background goroutine.
func StartServer(logger outbound.Logger, port int, serviceName string, isConnected func() bool) {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		status := HealthStatus{
			Status:      "ok",
			Connected:   isConnected(),
			ServiceName: serviceName,
		}

		if !status.Connected {
			status.Status = "degraded"
		}

		w.Header().Set("Content-Type", "application/json")

		if !status.Connected {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(status)
	})

	// TODO: Add a /metrics endpoint for Prometheus in a future sprint.

	addr := fmt.Sprintf(":%d", port)
	logger.Info("Starting internal HTTP server", "address", addr)

	go func() {
		if err := http.ListenAndServe(addr, mux); err != nil && err != http.ErrServerClosed {
			logger.Error("Internal HTTP server failed", err)
		}
	}()
}
