package fake

import (
	"context"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// FakeCodec is a mock implementation of the ProtoCodec interface.
type FakeCodec struct {
	logger outbound.Logger
}

// NewFakeCodec creates a new fake codec service.
func NewFakeCodec(logger outbound.Logger) outbound.ProtoCodec {
	return &FakeCodec{logger: logger}
}

// Encode logs the action and returns a dummy byte slice.
func (c *FakeCodec) Encode(ctx context.Context, v interface{}) ([]byte, error) {
	c.logger.Info("[FakeCodec] Simulating protobuf encoding")
	return []byte("encoded-proto-data"), nil
}

// Decode logs the action and does nothing to the target object.
func (c *FakeCodec) Decode(ctx context.Context, data []byte, v interface{}) error {
	c.logger.Info("[FakeCodec] Simulating protobuf decoding")
	// In a real scenario, we would unmarshal 'data' into 'v'.
	// Here, we do nothing, so 'v' remains unchanged.
	return nil
}

// Ensure FakeCodec implements the ProtoCodec interface.
var _ outbound.ProtoCodec = (*FakeCodec)(nil)
