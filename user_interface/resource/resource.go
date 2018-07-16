package resource

import "github.com/emicklei/go-restful"

type Resource interface {
	Register(container *restful.Container)
}
