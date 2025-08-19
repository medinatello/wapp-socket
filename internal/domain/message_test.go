package domain

import (
	"testing"
	"time"
)

func TestMessage(t *testing.T) {
	timestamp := time.Now()
	msg := Message{
		ID:        "msg-123",
		From:      JID("1234567890@s.whatsapp.net"),
		To:        JID("0987654321@s.whatsapp.net"),
		Text:      "Hello, World!",
		Timestamp: timestamp,
	}

	if msg.ID != "msg-123" {
		t.Errorf("Message.ID = %v, want msg-123", msg.ID)
	}

	if msg.From.String() != "1234567890@s.whatsapp.net" {
		t.Errorf("Message.From = %v, want 1234567890@s.whatsapp.net", msg.From)
	}

	if msg.To.String() != "0987654321@s.whatsapp.net" {
		t.Errorf("Message.To = %v, want 0987654321@s.whatsapp.net", msg.To)
	}

	if msg.Text != "Hello, World!" {
		t.Errorf("Message.Text = %v, want Hello, World!", msg.Text)
	}

	if !msg.Timestamp.Equal(timestamp) {
		t.Errorf("Message.Timestamp = %v, want %v", msg.Timestamp, timestamp)
	}
}

func TestMediaMessage(t *testing.T) {
	timestamp := time.Now()
	mediaMsg := MediaMessage{
		Message: Message{
			ID:        "media-123",
			From:      JID("1234567890@s.whatsapp.net"),
			To:        JID("0987654321@s.whatsapp.net"),
			Text:      "",
			Timestamp: timestamp,
		},
		MediaURL: "https://example.com/image.jpg",
		MimeType: "image/jpeg",
		Caption:  "A beautiful sunset",
	}

	if mediaMsg.ID != "media-123" {
		t.Errorf("MediaMessage.ID = %v, want media-123", mediaMsg.ID)
	}

	if mediaMsg.MediaURL != "https://example.com/image.jpg" {
		t.Errorf("MediaMessage.MediaURL = %v, want https://example.com/image.jpg", mediaMsg.MediaURL)
	}

	if mediaMsg.MimeType != "image/jpeg" {
		t.Errorf("MediaMessage.MimeType = %v, want image/jpeg", mediaMsg.MimeType)
	}

	if mediaMsg.Caption != "A beautiful sunset" {
		t.Errorf("MediaMessage.Caption = %v, want A beautiful sunset", mediaMsg.Caption)
	}
}

func TestEvent(t *testing.T) {
	message := &Message{
		ID:        "test-msg",
		From:      JID("sender@s.whatsapp.net"),
		To:        JID("recipient@s.whatsapp.net"),
		Text:      "Test message",
		Timestamp: time.Now(),
	}

	event := Event{
		Type:    EventMessageReceived,
		Payload: message,
	}

	if event.Type != EventMessageReceived {
		t.Errorf("Event.Type = %v, want %v", event.Type, EventMessageReceived)
	}

	payload, ok := event.Payload.(*Message)
	if !ok {
		t.Errorf("Event.Payload is not a *Message")
	}

	if payload.ID != "test-msg" {
		t.Errorf("Event.Payload.ID = %v, want test-msg", payload.ID)
	}
}

func TestEventConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{"EventMessageReceived", EventMessageReceived, "message_received"},
		{"EventPresenceUpdate", EventPresenceUpdate, "presence_update"},
		{"EventQRRefresh", EventQRRefresh, "qr_refresh"},
		{"EventDisconnected", EventDisconnected, "disconnected"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}
