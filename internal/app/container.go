package app

import (
	"os"

	slogadapter "github.com/medinatello/wapp-socket/internal/adapter/log/slog"
	storefake "github.com/medinatello/wapp-socket/internal/adapter/store/fake"
	wsfake "github.com/medinatello/wapp-socket/internal/adapter/ws/fake"
	"github.com/medinatello/wapp-socket/internal/port/outbound"
	"github.com/medinatello/wapp-socket/internal/telemetry"
	"github.com/medinatello/wapp-socket/internal/usecase"
)

// Container holds all the wired-up services for the application.
type Container struct {
	Config    *Config
	Logger    outbound.Logger
	Telemetry telemetry.Telemetry

	// Use Cases
	ConnectUseCase     *usecase.ConnectUseCase
	SendMessageUseCase *usecase.SendMessageUseCase
	ReceiveUseCase     *usecase.ReceiveUseCase
	GroupsUseCase      *usecase.GroupsUseCase
}

// NewContainer creates and wires up all the application components.
func NewContainer() (*Container, error) {
	// --- Configuration ---
	cfg, err := LoadConfig("./configs")
	if err != nil {
		return nil, err
	}

	// --- Logger ---
	logLevel := slogadapter.ParseLogLevel(cfg.App.LogLevel)
	logger := slogadapter.NewSlogLogger(os.Stdout, logLevel)

	// --- Telemetry ---
	tel := telemetry.NewOtelNoop()

	// --- Adapters (Outbound Ports) ---
	sessionAdapter := storefake.NewFakeStore(logger, cfg.Fakes.Seed)
	wsAdapter := wsfake.NewFakeWebSocketDialer(
		logger,
		cfg.Fakes.Seed,
		cfg.Fakes.ConnectTimeoutChance,
		cfg.Fakes.ConnectFailChance,
		cfg.Fakes.ReceiveIntervalMs,
	)

	// --- Use Cases ---
	connectUseCase := usecase.NewConnectUseCase(logger, wsAdapter, sessionAdapter)

	// The connProvider function is a closure that gets the active connection
	// from the ConnectUseCase. This is a simple way to share the connection state
	// for Sprint 1.
	connProvider := func() outbound.WebSocketConn {
		return connectUseCase.GetActiveConnection()
	}

	sendMessageUseCase := usecase.NewSendMessageUseCase(logger, connProvider)
	receiveUseCase := usecase.NewReceiveUseCase(logger, connProvider)
	groupsUseCase := usecase.NewGroupsUseCase(logger)

	// --- Build Container ---
	container := &Container{
		Config:             cfg,
		Logger:             logger,
		Telemetry:          tel,
		ConnectUseCase:     connectUseCase,
		SendMessageUseCase: sendMessageUseCase,
		ReceiveUseCase:     receiveUseCase,
		GroupsUseCase:      groupsUseCase,
	}

	logger.Info("Application container initialized successfully")
	return container, nil
}
