package middleware

import "github.com/emicklei/go-restful"

// Interface for middleware
type Middleware interface {
	Handle(req *restful.Request, resp *restful.Response, chain *restful.FilterChain)
}
