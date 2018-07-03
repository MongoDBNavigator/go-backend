package resource

import "github.com/emicklei/go-restful"

type ResourceInterface interface {
	Register(container *restful.Container)
}
