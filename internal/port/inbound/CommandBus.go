package inbound

import "context"

// Command is a placeholder for a generic command.
type Command interface {
	CommandType() string
}

// CommandHandler defines the interface for a handler that executes a command.
type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) (interface{}, error)
}

// CommandBus defines an interface for dispatching commands to their handlers.
// This is a placeholder for a potential future implementation.
type CommandBus interface {
	Dispatch(ctx context.Context, cmd Command) (interface{}, error)
	Register(cmdType string, handler CommandHandler) error
}
