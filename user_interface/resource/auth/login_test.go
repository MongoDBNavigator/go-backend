package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wsContainer := mux.NewRouter()

	username := "admin"
	password := "admin"
	expr := 24

	NewAuthResource(username, password, expr).Register(wsContainer)

	body := strings.NewReader(`{"username":"admin","password":"admin"}`)

	httpRequest, _ := http.NewRequest("POST", "http://localhost/api/v1/login", body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	wsContainer.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusOK, httpWriter.Code)
	assert.True(t, strings.Contains(httpWriter.Body.String(), "token"))
}

func TestLoginInvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wsContainer := mux.NewRouter()

	username := "admin"
	password := "admin"
	expr := 24

	NewAuthResource(username, password, expr).Register(wsContainer)

	body := strings.NewReader(`{"username":"admin1","password":"admin1"}`)

	httpRequest, _ := http.NewRequest("POST", "http://localhost/api/v1/login", body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	wsContainer.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpWriter.Code)

	assert.True(t, strings.Contains(httpWriter.Body.String(), "invalid credentials"))
}
