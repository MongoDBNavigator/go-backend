package swagger_resource

import (
	"github.com/MongoDBNavigator/go-backend/user_interface"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

// This resource generate json api docs for Swagger UI
type swaggerResource struct {
	sebServicesURL string
}

// Method to register resource
func (rcv *swaggerResource) Register(container *restful.Container) {
	container.Add(restfulspec.NewOpenAPIService(
		restfulspec.Config{
			WebServices:    container.RegisteredWebServices(), // you control what services are visible
			WebServicesURL: rcv.sebServicesURL,
			APIPath:        "/swagger/apidocs.json",
			PostBuildSwaggerObjectHandler: func(s *spec.Swagger) {
				s.Info = &spec.Info{
					InfoProps: spec.InfoProps{
						Title:       "MongoDB Navigator",
						Description: "MongoDB Navigator backend JSON API",
						Version:     "0.0.1",
					},
				}
			},
		},
	))
}

// Constructor for swaggerResource
func NewSwaggerResource(sebServicesURL string) user_interface.WebService {
	return &swaggerResource{
		sebServicesURL: sebServicesURL,
	}
}
