package slog

import (
	"bytes"
	"errors"
	"log/slog"
	"strings"
	"testing"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

func TestNewSlogLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelInfo)

	if logger == nil {
		t.Fatal("NewSlogLogger returned nil")
	}

	if logger.logger == nil {
		t.Fatal("NewSlogLogger created logger with nil internal logger")
	}
}

func TestSlogLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelDebug)

	logger.Debug("debug message", slog.String("key", "value"))

	output := buf.String()
	if !strings.Contains(output, "debug message") {
		t.Errorf("Debug log should contain message, got: %s", output)
	}
	if !strings.Contains(output, "DEBUG") {
		t.Errorf("Debug log should contain level DEBUG, got: %s", output)
	}
}

func TestSlogLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelInfo)

	logger.Info("info message", slog.String("key", "value"))

	output := buf.String()
	if !strings.Contains(output, "info message") {
		t.Errorf("Info log should contain message, got: %s", output)
	}
	if !strings.Contains(output, "INFO") {
		t.Errorf("Info log should contain level INFO, got: %s", output)
	}
}

func TestSlogLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelWarn)

	logger.Warn("warn message", slog.String("key", "value"))

	output := buf.String()
	if !strings.Contains(output, "warn message") {
		t.Errorf("Warn log should contain message, got: %s", output)
	}
	if !strings.Contains(output, "WARN") {
		t.Errorf("Warn log should contain level WARN, got: %s", output)
	}
}

func TestSlogLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelError)

	testErr := errors.New("test error")
	logger.Error("error message", testErr, slog.String("key", "value"))

	output := buf.String()
	if !strings.Contains(output, "error message") {
		t.Errorf("Error log should contain message, got: %s", output)
	}
	if !strings.Contains(output, "ERROR") {
		t.Errorf("Error log should contain level ERROR, got: %s", output)
	}
	if !strings.Contains(output, "test error") {
		t.Errorf("Error log should contain error message, got: %s", output)
	}
}

func TestSlogLogger_With(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelInfo)

	contextLogger := logger.With(slog.String("component", "test"))
	contextLogger.Info("test message")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("Context logger should contain message, got: %s", output)
	}
	if !strings.Contains(output, "component") {
		t.Errorf("Context logger should contain context key, got: %s", output)
	}
	if !strings.Contains(output, "test") {
		t.Errorf("Context logger should contain context value, got: %s", output)
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"error", slog.LevelError},
		{"unknown", slog.LevelInfo}, // Default case
		{"", slog.LevelInfo},        // Default case
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParseLogLevel(tt.input)
			if result.Level() != tt.expected {
				t.Errorf("ParseLogLevel(%s) = %v, want %v", tt.input, result.Level(), tt.expected)
			}
		})
	}
}

func TestSlogLogger_ImplementsInterface(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(&buf, slog.LevelInfo)

	// This test ensures that SlogLogger implements the outbound.Logger interface
	var _ outbound.Logger = logger
}
