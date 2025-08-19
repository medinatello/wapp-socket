package domain

import (
	"errors"
	"testing"
)

func TestDomainErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "ErrInvalidCredentials",
			err:      ErrInvalidCredentials,
			expected: "invalid credentials",
		},
		{
			name:     "ErrSessionNotFound",
			err:      ErrSessionNotFound,
			expected: "session not found",
		},
		{
			name:     "ErrConnectionFailed",
			err:      ErrConnectionFailed,
			expected: "connection failed",
		},
		{
			name:     "ErrNotConnected",
			err:      ErrNotConnected,
			expected: "not connected",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("%s.Error() = %v, want %v", tt.name, tt.err.Error(), tt.expected)
			}
		})
	}
}

func TestErrorsAreDistinct(t *testing.T) {
	// Ensure that each error is distinct and can be identified using errors.Is
	if errors.Is(ErrInvalidCredentials, ErrSessionNotFound) {
		t.Error("ErrInvalidCredentials should not be the same as ErrSessionNotFound")
	}

	if errors.Is(ErrConnectionFailed, ErrNotConnected) {
		t.Error("ErrConnectionFailed should not be the same as ErrNotConnected")
	}

	if errors.Is(ErrInvalidCredentials, ErrConnectionFailed) {
		t.Error("ErrInvalidCredentials should not be the same as ErrConnectionFailed")
	}
}

func TestErrorsIs(t *testing.T) {
	// Test that errors.Is works correctly with domain errors
	if !errors.Is(ErrInvalidCredentials, ErrInvalidCredentials) {
		t.Error("ErrInvalidCredentials should be equal to itself")
	}

	if !errors.Is(ErrSessionNotFound, ErrSessionNotFound) {
		t.Error("ErrSessionNotFound should be equal to itself")
	}
}
