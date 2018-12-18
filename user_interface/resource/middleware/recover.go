package middleware

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth/representation"
	"github.com/emicklei/go-restful"
)

type recoverMiddleware struct{}

// Recover and send 503
func (rcv *recoverMiddleware) Handle(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	defer func() {
		if r := recover(); r != nil {
			resp.WriteHeaderAndEntity(http.StatusServiceUnavailable, representation.Error{Message: "Try again later."})
		}
	}()

	chain.ProcessFilter(req, resp)
}

// Constructor for recoverMiddleware
func NewRecoverMiddleware() Middleware {
	return &recoverMiddleware{}
}
