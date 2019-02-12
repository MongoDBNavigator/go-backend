package middleware

import "net/http"

// Interface for middleware
type Middleware interface {
	Handle(next http.Handler) http.Handler
}
