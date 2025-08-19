package domain

import (
	"testing"
)

func TestJID_String(t *testing.T) {
	tests := []struct {
		name string
		jid  JID
		want string
	}{
		{
			name: "valid JID",
			jid:  JID("1234567890@s.whatsapp.net"),
			want: "1234567890@s.whatsapp.net",
		},
		{
			name: "empty JID",
			jid:  JID(""),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.jid.String(); got != tt.want {
				t.Errorf("JID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJID_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		jid  JID
		want bool
	}{
		{
			name: "empty JID",
			jid:  JID(""),
			want: true,
		},
		{
			name: "non-empty JID",
			jid:  JID("1234567890@s.whatsapp.net"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.jid.IsEmpty(); got != tt.want {
				t.Errorf("JID.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession(t *testing.T) {
	session := Session{
		ID:          "test-session-123",
		JID:         JID("1234567890@s.whatsapp.net"),
		IsConnected: true,
	}

	if session.ID != "test-session-123" {
		t.Errorf("Session.ID = %v, want test-session-123", session.ID)
	}

	if session.JID.String() != "1234567890@s.whatsapp.net" {
		t.Errorf("Session.JID = %v, want 1234567890@s.whatsapp.net", session.JID)
	}

	if !session.IsConnected {
		t.Errorf("Session.IsConnected = %v, want true", session.IsConnected)
	}
}
