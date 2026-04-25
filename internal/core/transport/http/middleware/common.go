package core_http_middleware

import (
	"github.com/google/uuid"
	"net/http"
)

func RequestId() Middleware {
	const requestIDHeader = "X-Request-Id"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Add(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}
