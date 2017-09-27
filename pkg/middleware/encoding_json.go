package middleware

import (
	"net/http"
)

// ContentTypeJSON sets response content type
// to application/json
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, req)
	})
}
