package fake

import (
	"bytes"
	"context"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// FakeCrypto is a mock implementation of the Crypto interface.
type FakeCrypto struct {
	logger outbound.Logger
}

// NewFakeCrypto creates a new fake crypto service.
func NewFakeCrypto(logger outbound.Logger) outbound.Crypto {
	return &FakeCrypto{logger: logger}
}

// Encrypt prepends "encrypted:" to the plaintext.
func (c *FakeCrypto) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	c.logger.Info("[FakeCrypto] Simulating encryption")
	return append([]byte("encrypted:"), plaintext...), nil
}

// Decrypt removes the "encrypted:" prefix.
func (c *FakeCrypto) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	c.logger.Info("[FakeCrypto] Simulating decryption")
	prefix := []byte("encrypted:")
	if !bytes.HasPrefix(ciphertext, prefix) {
		c.logger.Warn("[FakeCrypto] Ciphertext does not have expected prefix")
		return ciphertext, nil // Or return an error
	}
	return bytes.TrimPrefix(ciphertext, prefix), nil
}

// Ensure FakeCrypto implements the Crypto interface.
var _ outbound.Crypto = (*FakeCrypto)(nil)
