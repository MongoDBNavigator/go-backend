package auth

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"

	"github.com/MongoDBNavigator/go-backend/user_interface"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth/representation"
)

type authResource struct {
	username string
	password string
	exp      int
}

// Method to register resource
func (rcv *authResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/api/v1/login").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").
		To(rcv.login).
		Doc("Login (getting JWT token).").
		Reads(representation.PostCredentials{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.JwtToken{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusForbidden, http.StatusText(http.StatusForbidden), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{"Authorization"}))

	container.Add(ws)
}

// Constructor for swaggerResource
func NewAuthResource(username string, password string, exp int) user_interface.WebService {
	return &authResource{
		username: username,
		password: password,
		exp:      exp,
	}
}
