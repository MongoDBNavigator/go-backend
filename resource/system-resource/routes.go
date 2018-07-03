package system_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"github.com/MongoDBNavigator/go-backend/resource"
	"github.com/MongoDBNavigator/go-backend/resource/system-resource/representation"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

type systemResource struct {
	systemRepository repository.SystemRepositoryInterface
}

func (rcv *systemResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	tags := []string{"System info"}

	ws.Path("/api/v1/system").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/info").
		To(rcv.getInfo).
		Doc("Get system info (server version, etc.).").
		Writes(representation.Info{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Info{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}

func NewSystemResource(systemRepository repository.SystemRepositoryInterface) resource.ResourceInterface {
	return &systemResource{
		systemRepository: systemRepository,
	}
}
