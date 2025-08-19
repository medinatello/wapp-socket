package domain

// JID represents a Jabber ID, used to identify users and services.
// JID format follows the WhatsApp format: phone_number@s.whatsapp.net
type JID string

// String returns the string representation of the JID.
func (j JID) String() string {
	return string(j)
}

// IsEmpty checks if the JID is empty.
func (j JID) IsEmpty() bool {
	return j == ""
}

// Session represents the state of a connection.
// It contains the essential information needed to maintain a WhatsApp connection.
type Session struct {
	ID          string // Unique session identifier
	JID         JID    // Jabber ID of the connected user
	IsConnected bool   // Current connection status
	// In a real implementation, this would contain encryption keys, etc.
}
