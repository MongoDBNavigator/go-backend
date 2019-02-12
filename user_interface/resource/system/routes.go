package system

import (
	"github.com/gorilla/mux"

	"github.com/emicklei/go-restful"

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
func (rcv *systemResource) Register(container *mux.Router) {
	container.HandleFunc("/api/v1/system/info", rcv.getInfo).
		Methods("GET").
		Name("get_system_info")

	ws := new(restful.WebService)

	ws.Filter(rcv.jwtMiddleware.Handle)
	ws.Filter(rcv.recoverMiddleware.Handle)
	//
	//ws.Path("/api/v1/system").
	//	Consumes(restful.MIME_JSON).
	//	Produces(restful.MIME_JSON)
	//
	//ws.Route(ws.GET("/info").
	//	To(rcv.getInfo).
	//	Doc("Get system info (server version, etc.).").
	//	Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
	//	Writes(representation.Info{}).
	//	Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Info{}).
	//	Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
	//	Metadata(restfulspec.KeyOpenAPITags, []string{"System info"}))
	//
	//container.Add(ws)
}

// Constructor for systemResource
func NewSystemResource(
	systemInfoReader repository.SystemInfoReader,
	jwtMiddleware middleware.Middleware,
	recoverMiddleware middleware.Middleware,
) user_interface.WebService {
	return &systemResource{
		systemInfoReader:  systemInfoReader,
		jwtMiddleware:     jwtMiddleware,
		recoverMiddleware: recoverMiddleware,
	}
}
