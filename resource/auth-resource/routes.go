package auth_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource"
	"github.com/MongoDBNavigator/go-backend/resource/auth-resource/representation"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

type authResource struct {
	username string
	password string
	exp      int
}

func (rcv *authResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/api/v1/login").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").
		To(rcv.login).
		Doc("Login.").
		Reads(representation.PostCredentials{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.JwtToken{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusForbidden, http.StatusText(http.StatusForbidden), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{"Authorization"}))

	container.Add(ws)
}

func NewAuthResource(username string, password string, exp int) resource.ResourceInterface {
	return &authResource{
		username: username,
		password: password,
		exp:      exp,
	}
}
