package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// FakeMedia is a mock implementation of the MediaStore interface.
type FakeMedia struct {
	logger outbound.Logger
}

// NewFakeMedia creates a new fake media store.
func NewFakeMedia(logger outbound.Logger) outbound.MediaStore {
	return &FakeMedia{logger: logger}
}

// Save logs the action and returns a predictable, fake URL.
func (f *FakeMedia) Save(ctx context.Context, mediaData []byte, mimeType string) (string, error) {
	f.logger.Info("[FakeMedia] Simulating save", "size", len(mediaData), "mime", mimeType)
	// Generate a fake URL that looks plausible
	fakeURL := fmt.Sprintf("https://fake.media.store/uploads/%d.%s", time.Now().UnixNano(), mimeType)
	return fakeURL, nil
}

// Get logs the action and returns an empty byte slice.
func (f *FakeMedia) Get(ctx context.Context, url string) ([]byte, error) {
	f.logger.Info("[FakeMedia] Simulating get", "url", url)
	// Return empty data, as we don't actually store anything.
	return []byte{}, nil
}

// Ensure FakeMedia implements the MediaStore interface.
var _ outbound.MediaStore = (*FakeMedia)(nil)
