package outbound

// Logger defines a standard interface for structured logging.
// It supports structured logging with key-value pairs following slog conventions.
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, err error, args ...any)
}
