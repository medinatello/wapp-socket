package domain

import "errors"

// Domain-specific errors that can occur during WhatsApp operations.
var (
	// ErrInvalidCredentials indicates that the provided authentication credentials are invalid.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrSessionNotFound indicates that the requested session does not exist.
	ErrSessionNotFound = errors.New("session not found")

	// ErrConnectionFailed indicates that the connection attempt failed.
	ErrConnectionFailed = errors.New("connection failed")

	// ErrNotConnected indicates that an operation was attempted without an active connection.
	ErrNotConnected = errors.New("not connected")
)
