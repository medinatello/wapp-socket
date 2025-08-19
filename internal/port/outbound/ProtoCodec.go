package outbound

import "context"

// ProtoCodec defines the interface for encoding and decoding protobuf messages.
type ProtoCodec interface {
	Encode(ctx context.Context, v interface{}) ([]byte, error)
	Decode(ctx context.Context, data []byte, v interface{}) error
}
