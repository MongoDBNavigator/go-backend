package middleware

import "net/http"

type contentTypeMiddleware struct {
}

// Add header Content-Type
func (rcv *contentTypeMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

// Constructor for contentTypeMiddleware
func NewContentTypeMiddleware() Middleware {
	return &contentTypeMiddleware{}
}
