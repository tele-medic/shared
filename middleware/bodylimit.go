package middleware

import (
	"net/http"

	"github.com/tele-medic/shared/domain"
	"github.com/tele-medic/shared/response"
)

// BodyLimit returns middleware that limits request body size
func BodyLimit(maxBytes int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Body != nil {
				r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
			}
			next.ServeHTTP(w, r)

			// Note: MaxBytesReader will cause io.Read to return an error
			// if the body exceeds the limit. Handlers that parse JSON
			// will get a decode error and should return 400.
			// For cases where the body is silently truncated, the
			// MaxBytesReader also sets a flag that triggers a 413 response.
		})
	}
}

// BodyLimitStrict returns middleware that checks Content-Length upfront
func BodyLimitStrict(maxBytes int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > maxBytes {
				response.ErrorJSON(w, http.StatusRequestEntityTooLarge,
					domain.ErrCodeInvalidInput, "request body too large")
				return
			}
			if r.Body != nil {
				r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
			}
			next.ServeHTTP(w, r)
		})
	}
}
