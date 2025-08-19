package fake

import (
	"context"
	"math/rand"
	"time"

	"github.com/medinatello/wapp-socket/internal/domain"
	"github.com/medinatello/wapp-socket/internal/port/outbound"
)

// FakeWebSocketConn is a fake WebSocket connection that simulates receiving events.
type FakeWebSocketConn struct {
	logger   outbound.Logger
	events   chan domain.Event
	rand     *rand.Rand
	interval time.Duration
	stop     chan struct{}
}

// Send logs the frame and returns ok.
func (c *FakeWebSocketConn) Send(ctx context.Context, frame outbound.WebSocketFrame) error {
	c.logger.Info("[FakeWSConn] Sending frame", "frame", frame)
	// In a real scenario, we might simulate a delay here
	return nil
}

// Events returns a channel for receiving simulated events.
func (c *FakeWebSocketConn) Events() <-chan domain.Event {
	return c.events
}

// Close stops the event generation and closes the channel.
func (c *FakeWebSocketConn) Close() error {
	c.logger.Info("[FakeWSConn] Closing connection")
	close(c.stop) // Signal the goroutine to stop
	return nil
}

// start is a helper to begin the event generation loop.
func (c *FakeWebSocketConn) start() {
	go func() {
		defer close(c.events)
		ticker := time.NewTicker(c.interval)
		defer ticker.Stop()

		for {
			select {
			case <-c.stop:
				c.logger.Debug("[FakeWSConn] Stopping event generator")
				return
			case <-ticker.C:
				event := c.generateRandomEvent()
				c.logger.Debug("[FakeWSConn] Generating fake event", "type", event.Type)
				// Non-blocking send to prevent a slow consumer from blocking the generator
				select {
				case c.events <- event:
				default:
					c.logger.Warn("[FakeWSConn] Event channel is full. Dropping event.", "type", event.Type)
				}
			}
		}
	}()
}

func (c *FakeWebSocketConn) generateRandomEvent() domain.Event {
	eventTypes := []string{
		domain.EventMessageReceived,
		domain.EventPresenceUpdate,
		domain.EventQRRefresh,
		domain.EventDisconnected,
	}
	// Pick a random event type
	eventType := eventTypes[c.rand.Intn(len(eventTypes))]

	var payload interface{}
	switch eventType {
	case domain.EventMessageReceived:
		payload = &domain.Message{
			ID:        "fake-msg-" + time.Now().Format("150405"),
			From:      domain.JID("1234567890@s.whatsapp.net"),
			To:        domain.JID("0987654321@s.whatsapp.net"),
			Text:      "This is a simulated message.",
			Timestamp: time.Now(),
		}
	case domain.EventPresenceUpdate:
		payload = "composing" // A simple string for presence
	case domain.EventQRRefresh:
		payload = "new_qr_code_base64_string"
	case domain.EventDisconnected:
		payload = "connection closed by server"
	}

	return domain.Event{
		Type:    eventType,
		Payload: payload,
	}
}

// FakeWebSocketDialer is a fake WebSocket dialer that simulates connection outcomes.
type FakeWebSocketDialer struct {
	logger            outbound.Logger
	seed              int64
	timeoutChance     float64
	failChance        float64
	receiveIntervalMs int
}

// NewFakeWebSocketDialer creates a new fake dialer.
func NewFakeWebSocketDialer(logger outbound.Logger, seed int64, timeoutChance, failChance float64, interval int) outbound.WebSocketDialer {
	return &FakeWebSocketDialer{
		logger:            logger,
		seed:              seed,
		timeoutChance:     timeoutChance,
		failChance:        failChance,
		receiveIntervalMs: interval,
	}
}

// Dial simulates a connection attempt based on configured probabilities.
func (d *FakeWebSocketDialer) Dial(ctx context.Context, url string) (outbound.WebSocketConn, error) {
	d.logger.Info("[FakeWSDialer] Attempting to dial", "url", url)

	src := rand.NewSource(d.seed)
	if d.seed == 0 {
		src = rand.NewSource(time.Now().UnixNano())
	}
	r := rand.New(src)

	outcome := r.Float64()
	if outcome < d.failChance {
		d.logger.Warn("[FakeWSDialer] Simulating connection failure (invalid credentials)")
		return nil, domain.ErrInvalidCredentials
	}
	if outcome < d.failChance+d.timeoutChance {
		d.logger.Warn("[FakeWSDialer] Simulating connection timeout")
		// Simulate a delay for timeout
		time.Sleep(100 * time.Millisecond)
		return nil, context.DeadlineExceeded
	}

	d.logger.Info("[FakeWSDialer] Connection successful")
	conn := &FakeWebSocketConn{
		logger:   d.logger,
		events:   make(chan domain.Event, 10), // Buffered channel
		rand:     r,
		interval: time.Duration(d.receiveIntervalMs) * time.Millisecond,
		stop:     make(chan struct{}),
	}
	conn.start()

	return conn, nil
}

// Ensure the fakes implement the interfaces.
var _ outbound.WebSocketDialer = (*FakeWebSocketDialer)(nil)
var _ outbound.WebSocketConn = (*FakeWebSocketConn)(nil)
