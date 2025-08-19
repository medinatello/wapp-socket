package domain

// JID represents a Jabber ID, used to identify users and services.
type JID string

// Session represents the state of a connection.
type Session struct {
	ID         string
	JID        JID
	IsConnected bool
	// In a real implementation, this would contain encryption keys, etc.
}
