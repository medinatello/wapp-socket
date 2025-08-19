package usecase

import (
	"context"
	"fmt"

	"github.com/medinatello/wapp-socket/internal/domain"
	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// ReceiveUseCase handles the logic for processing incoming events from the WebSocket.
type ReceiveUseCase struct {
	logger       outbound.Logger
	connProvider func() outbound.WebSocketConn
}

// NewReceiveUseCase creates a new ReceiveUseCase.
func NewReceiveUseCase(logger outbound.Logger, connProvider func() outbound.WebSocketConn) *ReceiveUseCase {
	return &ReceiveUseCase{
		logger:       logger,
		connProvider: connProvider,
	}
}

// Execute starts a loop to listen for and process incoming events.
// It blocks until the context is done or the event channel is closed.
func (uc *ReceiveUseCase) Execute(ctx context.Context) error {
	uc.logger.Info("[ReceiveUseCase] Starting to listen for events")

	conn := uc.connProvider()
	if conn == nil {
		uc.logger.Error("[ReceiveUseCase] Cannot receive events, not connected", domain.ErrNotConnected)
		return domain.ErrNotConnected
	}

	eventCh := conn.Events()

	for {
		select {
		case <-ctx.Done():
			uc.logger.Info("[ReceiveUseCase] Stopping event listener due to context cancellation")
			return ctx.Err()
		case event, ok := <-eventCh:
			if !ok {
				uc.logger.Info("[ReceiveUseCase] Event channel closed, stopping listener")
				return nil
			}
			uc.processEvent(ctx, event)
		}
	}
}

func (uc *ReceiveUseCase) processEvent(ctx context.Context, event domain.Event) {
	uc.logger.Info("[ReceiveUseCase] Received event", "type", event.Type, "payload", fmt.Sprintf("%+v", event.Payload))
	// In a real application, this would involve more complex logic:
	// - Persisting messages
	// - Updating presence information
	// - Triggering other use cases (e.g., via an event bus)
}
