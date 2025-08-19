package usecase

import (
	"context"
	"fmt"

	"github.com/medinatello/wapp-socket/internal/domain"
	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// ConnectUseCase handles the logic for connecting to the WebSocket server.
type ConnectUseCase struct {
	logger       outbound.Logger
	wsDialer     outbound.WebSocketDialer
	sessionStore outbound.SessionStore
	// We need a place to store the active connection for other use cases.
	// This is a simplification for Sprint 1. A better approach might involve a connection manager.
	activeConn outbound.WebSocketConn
}

// NewConnectUseCase creates a new ConnectUseCase.
func NewConnectUseCase(logger outbound.Logger, wsDialer outbound.WebSocketDialer, sessionStore outbound.SessionStore) *ConnectUseCase {
	return &ConnectUseCase{
		logger:       logger,
		wsDialer:     wsDialer,
		sessionStore: sessionStore,
	}
}

// Execute attempts to establish a WebSocket connection and save the session.
func (uc *ConnectUseCase) Execute(ctx context.Context) (outbound.WebSocketConn, error) {
	uc.logger.Info("[ConnectUseCase] Executing connection sequence")

	// In a real app, URL would come from config or another service
	conn, err := uc.wsDialer.Dial(ctx, "wss://fake.whatsapp.server")
	if err != nil {
		uc.logger.Error("[ConnectUseCase] Failed to dial WebSocket", err)
		return nil, fmt.Errorf("dial failed: %w", err)
	}

	uc.logger.Info("[ConnectUseCase] Connection established, creating session")

	// For the fake implementation, we'll just create a dummy session
	session := &domain.Session{
		ID:         "fake-session-id",
		JID:        "1234567890@s.whatsapp.net",
		IsConnected: true,
	}

	if err := uc.sessionStore.Put(ctx, session); err != nil {
		uc.logger.Error("[ConnectUseCase] Failed to save session", err)
		conn.Close() // Clean up the connection
		return nil, fmt.Errorf("failed to save session: %w", err)
	}

	uc.logger.Info("[ConnectUseCase] Session saved successfully")
	uc.activeConn = conn // Store the active connection

	return conn, nil
}

// GetActiveConnection is a temporary method for Sprint 1 to allow other use cases
// to access the single, global connection.
func (uc *ConnectUseCase) GetActiveConnection() outbound.WebSocketConn {
	return uc.activeConn
}
