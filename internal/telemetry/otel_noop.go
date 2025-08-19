package telemetry

import "context"

// noopSpan is a mock span that does nothing.
type noopSpan struct{}

// End does nothing, fulfilling the Span interface.
func (s *noopSpan) End() {}

// OtelNoop is a no-op implementation of the Telemetry interface.
// It allows instrumenting the code without actually exporting any data.
type OtelNoop struct{}

// NewOtelNoop creates a new no-op telemetry provider.
func NewOtelNoop() *OtelNoop {
	return &OtelNoop{}
}

// StartSpan returns a no-op span and the original context. It does not record anything.
func (t *OtelNoop) StartSpan(ctx context.Context, name string, attrs ...Attribute) (context.Context, Span) {
	return ctx, &noopSpan{}
}

// RecordCounter does nothing. It does not record any metrics.
func (t *OtelNoop) RecordCounter(ctx context.Context, name string, value int64, attrs ...Attribute) {
	// No-op
}

// Ensure OtelNoop implements the Telemetry interface.
var _ Telemetry = (*OtelNoop)(nil)
