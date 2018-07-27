package user_interface

import "github.com/emicklei/go-restful"

type WebService interface {
	Register(container *restful.Container)
}
