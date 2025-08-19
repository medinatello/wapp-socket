package slog

import (
	"io"
	"log/slog"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// SlogLogger is an implementation of the Logger interface using the standard log/slog package.
type SlogLogger struct {
	logger *slog.Logger
}

// NewSlogLogger creates a new logger that writes to the provided io.Writer.
func NewSlogLogger(w io.Writer, level slog.Leveler) *SlogLogger {
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: level,
	})
	return &SlogLogger{
		logger: slog.New(handler),
	}
}

// Debug logs a message at the debug level.
func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs a message at the info level.
func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a message at the warning level.
func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs a message at the error level, including the error.
func (l *SlogLogger) Error(msg string, err error, args ...any) {
	// Combine the error and other args for structured logging.
	allArgs := append(args, slog.String("error", err.Error()))
	l.logger.Error(msg, allArgs...)
}

// The following methods are not part of the interface but are useful for slog.
// With converts the logger to a structured logger with the given attributes.
func (l *SlogLogger) With(args ...any) outbound.Logger {
	return &SlogLogger{
		logger: l.logger.With(args...),
	}
}

// To satisfy slog.Handler interface if needed, but we are wrapping it.
// We need to make sure our logger implements our interface.
var _ outbound.Logger = (*SlogLogger)(nil)

// ParseLogLevel converts a string level to a slog.Level.
func ParseLogLevel(levelStr string) slog.Leveler {
	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
