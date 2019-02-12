package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth/representation"
	"github.com/dgrijalva/jwt-go"
)

// Method to get JWT token
func (rcv *authResource) login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var postRequest representation.PostCredentials

	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if postRequest.Username != rcv.username || postRequest.Password != rcv.password {
		rcv.writeErrorResponse(w, http.StatusForbidden, errors.New("invalid credentials"))
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = rcv.username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(rcv.exp)).Unix()

	tokenString, err := token.SignedString([]byte(rcv.password))

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	result := new(representation.JwtToken)
	result.Token = tokenString

	rcv.writeResponse(w, http.StatusOK, result)
}
