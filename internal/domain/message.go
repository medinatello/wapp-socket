package domain

import "time"

// Message represents a basic text message.
type Message struct {
	ID        string
	From      JID
	To        JID
	Text      string
	Timestamp time.Time
}

// MediaMessage represents a message with media content.
type MediaMessage struct {
	Message
	MediaURL string
	MimeType string
	Caption  string
}

// Event represents a generic event from the WebSocket stream.
type Event struct {
	Type    string
	Payload interface{} // Can be a Message, Presence, etc.
}

const (
	EventMessageReceived = "message_received"
	EventPresenceUpdate  = "presence_update"
	EventQRRefresh       = "qr_refresh"
	EventDisconnected    = "disconnected"
)
