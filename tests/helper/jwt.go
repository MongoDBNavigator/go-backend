package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const PASSWORD = "admin"

func GenerateJwtToken() string {
	token := jwt.New(jwt.SigningMethodHS256)

	username := "admin"
	password := PASSWORD

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()

	tokenString, _ := token.SignedString([]byte(password))

	return tokenString
}
