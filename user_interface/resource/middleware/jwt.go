package middleware

import (
	"log"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

type jwtMiddleware struct {
	password string
}

// Check JWT token
func (rcv *jwtMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		if err := jwtMiddleware.CheckJWT(w, r); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	})
}

// Constructor for jwtMiddleware
func NewJwtMiddleware(password string) Middleware {
	return &jwtMiddleware{
		password: password,
	}
}
