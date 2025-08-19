package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/medinatello/wapp-socket/internal/app"
	"github.com/medinatello/wapp-socket/interface/http"
)

func main() {
	// --- Initialization ---
	container, err := app.NewContainer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize application: %v\n", err)
		os.Exit(1)
	}
	logger := container.Logger
	logger.Info("Starting whatsd daemon...")

	// --- Connect ---
	logger.Info("Attempting to connect...")
	_, err = container.ConnectUseCase.Execute(context.Background())
	if err != nil {
		logger.Error("Failed to connect on startup, shutting down.", err)
		os.Exit(1)
	}
	logger.Info("Successfully connected.")

	// --- Start HTTP Server ---
	// The isConnected function provides live status to the health check.
	isConnected := func() bool {
		conn := container.ConnectUseCase.GetActiveConnection()
		return conn != nil
	}
	// A real app would get the port from config. Hardcoding for now.
	http.StartServer(logger, 8081, container.Config.App.Name, isConnected)

	// --- Start Event Stream ---
	// Create a context that can be cancelled for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		logger.Info("Starting event stream processing...")
		if err := container.ReceiveUseCase.Execute(ctx); err != nil {
			logger.Error("Event stream failed", err)
			// In a real app, you might want to trigger a reconnect or shutdown.
		}
		logger.Info("Event stream stopped.")
	}()

	// --- Graceful Shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // Block until a signal is received

	logger.Info("Shutting down daemon...")

	// Signal the event stream to stop
	cancel()

	// Close the WebSocket connection
	if conn := container.ConnectUseCase.GetActiveConnection(); conn != nil {
		conn.Close()
	}

	logger.Info("Daemon shut down gracefully.")
}
