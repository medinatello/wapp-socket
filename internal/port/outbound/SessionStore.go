package outbound

import (
	"context"

	"github.com/medinatello/wapp-socket/internal/domain"
)

// SessionStore defines the interface for persisting and retrieving session data.
type SessionStore interface {
	Get(ctx context.Context, jid domain.JID) (*domain.Session, error)
	Put(ctx context.Context, session *domain.Session) error
	Delete(ctx context.Context, jid domain.JID) error
}
