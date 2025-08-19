package fake

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/medinatello/wapp-socket/internal/domain"
	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// FakeStore is an in-memory implementation of the SessionStore interface.
// It includes small random delays to simulate network latency.
type FakeStore struct {
	logger outbound.Logger
	mu     sync.RWMutex
	data   map[domain.JID]*domain.Session
	rand   *rand.Rand
}

// NewFakeStore creates a new in-memory store.
func NewFakeStore(logger outbound.Logger, seed int64) outbound.SessionStore {
	src := rand.NewSource(seed)
	if seed == 0 {
		src = rand.NewSource(time.Now().UnixNano())
	}
	return &FakeStore{
		logger: logger,
		data:   make(map[domain.JID]*domain.Session),
		rand:   rand.New(src),
	}
}

// simulateLatency introduces a small, random delay to mock I/O operations.
func (s *FakeStore) simulateLatency() {
	// Simulate a small delay between 10ms and 50ms
	delay := time.Duration(s.rand.Intn(41)+10) * time.Millisecond
	time.Sleep(delay)
}

// Get retrieves a session from the in-memory map.
func (s *FakeStore) Get(ctx context.Context, jid domain.JID) (*domain.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.simulateLatency()

	session, ok := s.data[jid]
	if !ok {
		s.logger.Warn("[FakeStore] Session not found", "jid", jid)
		return nil, domain.ErrSessionNotFound
	}

	s.logger.Info("[FakeStore] Session retrieved", "jid", jid)
	return session, nil
}

// Put adds or updates a session in the in-memory map.
func (s *FakeStore) Put(ctx context.Context, session *domain.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.simulateLatency()

	if session == nil {
		return nil // Or return an error
	}
	s.data[session.JID] = session
	s.logger.Info("[FakeStore] Session saved", "jid", session.JID)
	return nil
}

// Delete removes a session from the in-memory map.
func (s *FakeStore) Delete(ctx context.Context, jid domain.JID) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.simulateLatency()

	delete(s.data, jid)
	s.logger.Info("[FakeStore] Session deleted", "jid", jid)
	return nil
}

// Ensure FakeStore implements the SessionStore interface.
var _ outbound.SessionStore = (*FakeStore)(nil)
