package response

import (
	"encoding/json"
	"net/http"
)

// Response is the standard API response wrapper
type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// Error represents an API error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// PaginatedResponse wraps paginated data
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       any         `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// Pagination contains pagination metadata
type Pagination struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasMore    bool `json:"has_more"`
}

// JSON writes a JSON response
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success: status >= 200 && status < 300,
		Data:    data,
	})
}

// ErrorJSON writes an error JSON response
func ErrorJSON(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	})
}

// PaginatedJSON writes a paginated JSON response
func PaginatedJSON(w http.ResponseWriter, data any, page, perPage, total int) {
	totalPages := total / perPage
	if total%perPage != 0 {
		totalPages++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PaginatedResponse{
		Success: true,
		Data:    data,
		Pagination: &Pagination{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: totalPages,
			HasMore:    page < totalPages,
		},
	})
}
