package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MongoDBNavigator/go-backend/user_interface"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth/representation"
)

type authResource struct {
	username string
	password string
	exp      int
}

// write status & body to response
func (rcv *authResource) writeResponse(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err)
	}
}

// write status & body to response
func (rcv *authResource) writeErrorResponse(w http.ResponseWriter, status int, err error) {
	rcv.writeResponse(w, status, representation.Error{Message: err.Error()})
}

// Method to register resource
func (rcv *authResource) Register(container *mux.Router) {
	container.HandleFunc("/api/v1/login", rcv.login).
		Methods("POST").
		Name("login")
}

// Constructor for swaggerResource
func NewAuthResource(username string, password string, exp int) user_interface.WebService {
	return &authResource{
		username: username,
		password: password,
		exp:      exp,
	}
}
