package system_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"

	"github.com/MongoDBNavigator/go-backend/domain/system/repository"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system_resource/representation"
)

type systemResource struct {
	systemInfoReader repository.SystemInfoReader
	jwtMiddleware    middleware.Middleware
}

// Method to register resource
func (rcv *systemResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Filter(rcv.jwtMiddleware.Handle)

	ws.Path("/api/v1/system").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/info").
		To(rcv.getInfo).
		Doc("Get system info (server version, etc.).").
		Writes(representation.Info{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Info{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{"System info"}))

	container.Add(ws)
}

// Constructor for systemResource
func NewSystemResource(systemInfoReader repository.SystemInfoReader, jwtMiddleware middleware.Middleware) resource.Resource {
	return &systemResource{
		systemInfoReader: systemInfoReader,
		jwtMiddleware:    jwtMiddleware,
	}
}
