package swagger_resource

import (
	"github.com/MongoDBNavigator/go-backend/resource"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

type swaggerResource struct {
	sebServicesURL string
}

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

func NewSwaggerResource(sebServicesURL string) resource.ResourceInterface {
	return &swaggerResource{
		sebServicesURL: sebServicesURL,
	}
}
