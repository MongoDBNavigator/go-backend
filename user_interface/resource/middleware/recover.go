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
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				if err := json.NewEncoder(w).Encode(representation.Error{Message: "Try again later."}); err != nil {
					log.Println(err)
				}
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Constructor for recoverMiddleware
func NewRecoverMiddleware() Middleware {
	return &recoverMiddleware{}
}
