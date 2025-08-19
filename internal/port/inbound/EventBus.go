package inbound

import (
	"context"

	"github.com/medinatello/wapp-socket/internal/domain"
)

// EventBus defines an interface for a publish-subscribe mechanism for domain events.
type EventBus interface {
	Publish(ctx context.Context, topic string, event domain.Event) error
	Subscribe(ctx context.Context, topic string) (<-chan domain.Event, error)
}
