package domain

import (
	"time"
)

// SessionStatus represents the state of a session
type SessionStatus string

const (
	SessionStatusWaiting   SessionStatus = "waiting"
	SessionStatusActive    SessionStatus = "active"
	SessionStatusEnded     SessionStatus = "ended"
	SessionStatusCancelled SessionStatus = "cancelled"
)

// ParticipantRole represents the role in a session
type ParticipantRole string

const (
	RolePatient  ParticipantRole = "patient"
	RoleNurse    ParticipantRole = "nurse"
	RoleDoctor   ParticipantRole = "doctor"
	RoleObserver ParticipantRole = "observer"
)

// ParticipantStatus represents connection state
type ParticipantStatus string

const (
	ParticipantStatusInvited      ParticipantStatus = "invited"
	ParticipantStatusWaiting      ParticipantStatus = "waiting"
	ParticipantStatusConnected    ParticipantStatus = "connected"
	ParticipantStatusDisconnected ParticipantStatus = "disconnected"
)

// QueueStatus represents waiting queue state
type QueueStatus string

const (
	QueueStatusWaiting   QueueStatus = "waiting"
	QueueStatusAssigned  QueueStatus = "assigned"
	QueueStatusCompleted QueueStatus = "completed"
	QueueStatusExpired   QueueStatus = "expired"
)

// SessionType represents the type of consultation
type SessionType string

const (
	SessionTypeVideo    SessionType = "video"
	SessionTypeAudio    SessionType = "audio_only"
	SessionTypeChatOnly SessionType = "chat_only"
)

// Session represents a consultation session
type Session struct {
	ID              string            `json:"id"`
	TenantID        string            `json:"tenant_id"`
	RoomName        string            `json:"room_name"`
	Status          SessionStatus     `json:"status"`
	Type            SessionType       `json:"type"`
	MaxParticipants int               `json:"max_participants"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	ScheduledAt     *time.Time        `json:"scheduled_at,omitempty"`
	StartedAt       *time.Time        `json:"started_at,omitempty"`
	EndedAt         *time.Time        `json:"ended_at,omitempty"`
	DurationSecs    int               `json:"duration_secs,omitempty"`
	EndReason       string            `json:"end_reason,omitempty"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

// Participant represents a user in a session
type Participant struct {
	ID           string            `json:"id"`
	SessionID    string            `json:"session_id"`
	Identity     string            `json:"identity"`
	Name         string            `json:"name"`
	Role         ParticipantRole   `json:"role"`
	Status       ParticipantStatus `json:"status"`
	Permissions  map[string]bool   `json:"permissions"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	JoinedAt     *time.Time        `json:"joined_at,omitempty"`
	LeftAt       *time.Time        `json:"left_at,omitempty"`
	DurationSecs int              `json:"duration_secs,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
}

// QueueEntry represents a patient waiting in queue
type QueueEntry struct {
	ID              string      `json:"id"`
	TenantID        string      `json:"tenant_id"`
	SessionID       string      `json:"session_id"`
	PatientIdentity string      `json:"patient_identity"`
	Priority        int         `json:"priority"`
	Position        int         `json:"position"`
	Status          QueueStatus `json:"status"`
	AssignedTo      string      `json:"assigned_to,omitempty"`
	QueuedAt        time.Time   `json:"queued_at"`
	AssignedAt      *time.Time  `json:"assigned_at,omitempty"`
	CompletedAt     *time.Time  `json:"completed_at,omitempty"`
}

// Tenant represents an organization using the platform
type Tenant struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Domain    string         `json:"domain,omitempty"`
	Plan      string         `json:"plan"`
	IsActive  bool           `json:"is_active"`
	Settings  map[string]any `json:"settings,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// APIKey represents an API key for a tenant
type APIKey struct {
	ID        string     `json:"id"`
	TenantID  string     `json:"tenant_id"`
	KeyPrefix string     `json:"key_prefix"`
	KeyHash   string     `json:"-"`
	Name      string     `json:"name"`
	Scopes    []string   `json:"scopes"`
	IsActive  bool       `json:"is_active"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	LastUsed  *time.Time `json:"last_used,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// ChatMessage represents an in-call text message
type ChatMessage struct {
	ID             string            `json:"id"`
	SessionID      string            `json:"session_id"`
	SenderIdentity string            `json:"sender_identity"`
	MessageType    string            `json:"message_type"`
	Content        string            `json:"content"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
}

// Recording represents a call recording
type Recording struct {
	ID            string     `json:"id"`
	SessionID     string     `json:"session_id"`
	TenantID      string     `json:"tenant_id"`
	EgressID      string     `json:"egress_id,omitempty"`
	Status        string     `json:"status"`
	RecordingType string     `json:"recording_type"`
	StoragePath   string     `json:"storage_path,omitempty"`
	FileSizeBytes int64      `json:"file_size_bytes,omitempty"`
	DurationSecs  int        `json:"duration_secs,omitempty"`
	Format        string     `json:"format"`
	ConsentGiven  bool       `json:"consent_given"`
	ConsentBy     []string   `json:"consent_by"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// WebhookEndpoint represents a tenant's webhook configuration
type WebhookEndpoint struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	URL       string    `json:"url"`
	Secret    string    `json:"-"`
	Events    []string  `json:"events"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
