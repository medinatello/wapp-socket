package telemetry

import "context"

// Attribute is a key-value pair used for telemetry.
type Attribute struct {
	Key   string
	Value interface{}
}

// Span is a placeholder for a tracing span.
type Span interface {
	End()
}

// Telemetry defines a standard interface for telemetry (tracing and metrics).
// This provides a stable internal API that can be backed by OpenTelemetry or a no-op implementation.
type Telemetry interface {
	StartSpan(ctx context.Context, name string, attrs ...Attribute) (context.Context, Span)
	RecordCounter(ctx context.Context, name string, value int64, attrs ...Attribute)
}
