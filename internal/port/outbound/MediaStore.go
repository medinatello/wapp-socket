package outbound

import "context"

// MediaStore defines the interface for saving and retrieving media files.
type MediaStore interface {
	Save(ctx context.Context, mediaData []byte, mimeType string) (url string, err error)
	Get(ctx context.Context, url string) (mediaData []byte, err error)
}
