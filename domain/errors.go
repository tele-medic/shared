package domain

import "fmt"

// Error codes
const (
	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeAlreadyExists = "ALREADY_EXISTS"
	ErrCodeInvalidInput  = "INVALID_INPUT"
	ErrCodeUnauthorized  = "UNAUTHORIZED"
	ErrCodeForbidden     = "FORBIDDEN"
	ErrCodeInternal      = "INTERNAL"
	ErrCodeRoomFull      = "ROOM_FULL"
	ErrCodeSessionEnded  = "SESSION_ENDED"
	ErrCodeQueueFull     = "QUEUE_FULL"
	ErrCodeRateLimited   = "RATE_LIMITED"
)

// DomainError represents a structured error
type DomainError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

func NewNotFoundError(resource string) *DomainError {
	return &DomainError{Code: ErrCodeNotFound, Message: fmt.Sprintf("%s not found", resource)}
}

func NewAlreadyExistsError(resource string) *DomainError {
	return &DomainError{Code: ErrCodeAlreadyExists, Message: fmt.Sprintf("%s already exists", resource)}
}

func NewInvalidInputError(msg string) *DomainError {
	return &DomainError{Code: ErrCodeInvalidInput, Message: msg}
}

func NewUnauthorizedError(msg string) *DomainError {
	return &DomainError{Code: ErrCodeUnauthorized, Message: msg}
}

func NewForbiddenError(msg string) *DomainError {
	return &DomainError{Code: ErrCodeForbidden, Message: msg}
}

func NewInternalError(msg string, err error) *DomainError {
	return &DomainError{Code: ErrCodeInternal, Message: msg, Err: err}
}

func NewRoomFullError() *DomainError {
	return &DomainError{Code: ErrCodeRoomFull, Message: "room has reached maximum participants"}
}

func NewSessionEndedError() *DomainError {
	return &DomainError{Code: ErrCodeSessionEnded, Message: "session has already ended"}
}

// IsNotFound checks if error is a not found error
func IsNotFound(err error) bool {
	if e, ok := err.(*DomainError); ok {
		return e.Code == ErrCodeNotFound
	}
	return false
}

// IsAlreadyExists checks if error is an already exists error
func IsAlreadyExists(err error) bool {
	if e, ok := err.(*DomainError); ok {
		return e.Code == ErrCodeAlreadyExists
	}
	return false
}
