package outbound

import (
	"context"

	"github.com/medinatello/wapp-socket/internal/domain"
)

// WebSocketFrame represents a generic WebSocket message frame.
// In a real implementation, this would be more specific (e.g., []byte).
type WebSocketFrame interface{}

// WebSocketConn defines the interface for an active WebSocket connection.
type WebSocketConn interface {
	Send(ctx context.Context, frame WebSocketFrame) error
	Events() <-chan domain.Event
	Close() error
}

// WebSocketDialer defines the interface for creating a WebSocket connection.
type WebSocketDialer interface {
	Dial(ctx context.Context, url string) (WebSocketConn, error)
}
