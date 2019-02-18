package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth/representation"
)

type recoverMiddleware struct{}

// Recover and send 503
func (rcv *recoverMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		if err := json.NewEncoder(w).Encode(representation.Error{Message: "Try again later."}); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	})
}

// Constructor for recoverMiddleware
func NewRecoverMiddleware() Middleware {
	return &recoverMiddleware{}
}
