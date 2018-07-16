package middleware

import (
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
)

type jwtMiddleware struct {
	password string
}

// Check JWT token
func (rcv *jwtMiddleware) Handle(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(rcv.password), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			switch err {
			case "token contains an invalid number of segments":
				http.Error(w, err, http.StatusForbidden)
			case "signature is invalid":
				http.Error(w, err, http.StatusForbidden)
			default:
				http.Error(w, err, http.StatusUnauthorized)
			}
		},
	})

	if err := jwtMiddleware.CheckJWT(resp.ResponseWriter, req.Request); err != nil {
		return
	}

	chain.ProcessFilter(req, resp)
}

// Constructor for jwtMiddleware
func NewJwtMiddleware(password string) Middleware {
	return &jwtMiddleware{
		password: password,
	}
}
