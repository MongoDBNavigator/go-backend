package system

import (
	"github.com/gorilla/mux"

	"github.com/MongoDBNavigator/go-backend/domain/system/repository"
	"github.com/MongoDBNavigator/go-backend/user_interface"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
)

type systemResource struct {
	systemInfoReader  repository.SystemInfoReader
	jwtMiddleware     middleware.Middleware
	recoverMiddleware middleware.Middleware
}

// Method to register resource
func (rcv *systemResource) Register(r *mux.Router) {
	sr := r.PathPrefix("/api/v1/system").Subrouter()
	sr.Use(rcv.jwtMiddleware.Handle)

	sr.HandleFunc("/info", rcv.getInfo).
		Methods("GET").
		Name("get_system_info")

}

// Constructor for systemResource
func NewSystemResource(
	systemInfoReader repository.SystemInfoReader,
	jwtMiddleware middleware.Middleware,
) user_interface.WebService {
	return &systemResource{
		systemInfoReader: systemInfoReader,
		jwtMiddleware:    jwtMiddleware,
	}
}
