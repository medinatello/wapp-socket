package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrSessionNotFound    = errors.New("session not found")
	ErrConnectionFailed   = errors.New("connection failed")
	ErrNotConnected       = errors.New("not connected")
)
