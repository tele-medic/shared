package domain

import "fmt"

// Validation constants
const (
	MaxNameLength     = 255
	MaxIdentityLength = 255
	MaxDomainLength   = 255
	MaxChatContent    = 10000
	MaxParticipants   = 100
)

// Valid plans for tenants
var validPlans = map[string]bool{
	"free":       true,
	"starter":    true,
	"pro":        true,
	"enterprise": true,
}

// Valid scopes for API keys
var validScopes = map[string]bool{
	"*":              true,
	"sessions:read":  true,
	"sessions:write": true,
	"queue:read":     true,
	"queue:write":    true,
	"recordings:read": true,
	"recordings:write": true,
	"webhooks:read":  true,
	"webhooks:write": true,
}

// IsValidSessionType checks if the session type is valid
func IsValidSessionType(t string) bool {
	switch SessionType(t) {
	case SessionTypeVideo, SessionTypeAudio, SessionTypeChatOnly:
		return true
	}
	return false
}

// IsValidParticipantRole checks if the role is valid
func IsValidParticipantRole(r string) bool {
	switch ParticipantRole(r) {
	case RolePatient, RoleNurse, RoleDoctor, RoleObserver:
		return true
	}
	return false
}

// IsValidPlan checks if the plan is valid
func IsValidPlan(plan string) bool {
	return validPlans[plan]
}

// ValidateScopes checks that all scopes are valid, returns first invalid scope
func ValidateScopes(scopes []string) error {
	if len(scopes) == 0 {
		return fmt.Errorf("at least one scope is required")
	}
	for _, s := range scopes {
		if !validScopes[s] {
			return fmt.Errorf("invalid scope: %s", s)
		}
	}
	return nil
}

// ValidateStringLength checks that a string does not exceed max length
func ValidateStringLength(field, value string, maxLen int) error {
	if len(value) > maxLen {
		return fmt.Errorf("%s exceeds maximum length of %d", field, maxLen)
	}
	return nil
}
