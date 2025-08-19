package outbound

import "context"

// Crypto defines the interface for cryptographic operations.
type Crypto interface {
	Encrypt(ctx context.Context, plaintext []byte) (ciphertext []byte, err error)
	Decrypt(ctx context.Context, ciphertext []byte) (plaintext []byte, err error)
}
