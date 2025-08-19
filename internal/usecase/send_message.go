package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/medinatello/wapp-socket/internal/domain"
	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// SendMessageUseCase handles the logic for sending a text message.
type SendMessageUseCase struct {
	logger     outbound.Logger
	// This is a simplification for Sprint 1. A better way is to get the conn from a shared manager.
	connProvider func() outbound.WebSocketConn
}

// NewSendMessageUseCase creates a new SendMessageUseCase.
func NewSendMessageUseCase(logger outbound.Logger, connProvider func() outbound.WebSocketConn) *SendMessageUseCase {
	return &SendMessageUseCase{
		logger:     logger,
		connProvider: connProvider,
	}
}

// Execute sends a text message to a given JID.
func (uc *SendMessageUseCase) Execute(ctx context.Context, to domain.JID, text string) error {
	uc.logger.Info("[SendMessageUseCase] Executing send message", "to", to)

	conn := uc.connProvider()
	if conn == nil {
		uc.logger.Error("[SendMessageUseCase] Cannot send message, not connected", domain.ErrNotConnected)
		return domain.ErrNotConnected
	}

	// In a real implementation, the frame would be a complex protobuf object.
	// Here, we just send the domain message itself as the "frame".
	messageFrame := &domain.Message{
		ID:        fmt.Sprintf("msg_%d", time.Now().UnixNano()),
		To:        to,
		Text:      text,
		Timestamp: time.Now(),
	}

	if err := conn.Send(ctx, messageFrame); err != nil {
		uc.logger.Error("[SendMessageUseCase] Failed to send message frame", err)
		return fmt.Errorf("send failed: %w", err)
	}

	uc.logger.Info("[SendMessageUseCase] Message sent successfully", "to", to, "id", messageFrame.ID)
	return nil
}
