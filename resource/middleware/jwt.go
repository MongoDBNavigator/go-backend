package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
)

type jwtMiddleware struct {
	password string
}

func (rcv *jwtMiddleware) Handle(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(rcv.password), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	if err := jwtMiddleware.CheckJWT(resp.ResponseWriter, req.Request); err != nil {
		return
	}

	chain.ProcessFilter(req, resp)
}

func NewJwtMiddleware(password string) Middleware {
	return &jwtMiddleware{
		password: password,
	}
}
