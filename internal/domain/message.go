package domain

import "time"

// Message represents a basic text message in the WhatsApp protocol.
// It contains all the essential information for a text-based communication.
type Message struct {
	ID        string    // Unique message identifier
	From      JID       // Sender's Jabber ID
	To        JID       // Recipient's Jabber ID
	Text      string    // Message content
	Timestamp time.Time // When the message was sent/received
}

// MediaMessage represents a message with media content.
// It embeds Message and adds media-specific fields.
type MediaMessage struct {
	Message
	MediaURL string // URL or path to the media file
	MimeType string // MIME type of the media (image/jpeg, video/mp4, etc.)
	Caption  string // Optional caption for the media
}

// Event represents a generic event from the WebSocket stream.
// Events are used to communicate different types of updates from the WhatsApp connection.
type Event struct {
	Type    string // Event type identifier (see constants below)
	Payload any    // Event-specific data (can be a Message, Presence, etc.)
}

// Event type constants define the different types of events that can be received.
const (
	// EventMessageReceived indicates a new message was received
	EventMessageReceived = "message_received"
	// EventPresenceUpdate indicates a contact's presence status changed
	EventPresenceUpdate = "presence_update"
	// EventQRRefresh indicates the QR code for authentication needs to be refreshed
	EventQRRefresh = "qr_refresh"
	// EventDisconnected indicates the connection was lost
	EventDisconnected = "disconnected"
)
