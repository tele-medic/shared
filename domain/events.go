package domain

import "time"

// EventType represents the type of domain event
type EventType string

const (
	EventSessionCreated   EventType = "session.created"
	EventSessionStarted   EventType = "session.started"
	EventSessionEnded     EventType = "session.ended"
	EventSessionCancelled EventType = "session.cancelled"

	EventParticipantJoined  EventType = "participant.joined"
	EventParticipantLeft    EventType = "participant.left"
	EventParticipantMuted   EventType = "participant.muted"
	EventParticipantUnmuted EventType = "participant.unmuted"

	EventQueueEntryAdded    EventType = "queue.entry_added"
	EventQueueEntryAssigned EventType = "queue.entry_assigned"
	EventQueueEntryExpired  EventType = "queue.entry_expired"

	EventRecordingStarted EventType = "recording.started"
	EventRecordingStopped EventType = "recording.stopped"
	EventRecordingReady   EventType = "recording.ready"
	EventRecordingFailed  EventType = "recording.failed"

	EventChatMessageSent EventType = "chat.message_sent"
)

// Event represents a domain event for inter-service communication
type Event struct {
	ID        string         `json:"id"`
	Type      EventType      `json:"type"`
	TenantID  string         `json:"tenant_id"`
	SessionID string         `json:"session_id,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
}
