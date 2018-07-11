package middleware

import "github.com/emicklei/go-restful"

type Middleware interface {
	Handle(req *restful.Request, resp *restful.Response, chain *restful.FilterChain)
}
