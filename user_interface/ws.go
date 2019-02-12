package user_interface

import (
	"github.com/gorilla/mux"
)

type WebService interface {
	Register(container *mux.Router)
}
