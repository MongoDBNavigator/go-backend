package database

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"errors"

	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/stretchr/testify/assert"
)

func TestPostDatabaseSuccess(t *testing.T) {
	container := initResource(t)

	databaseJson := `{"name":"MyDB"}`

	var database *representation.PostDatabase

	json.Unmarshal([]byte(databaseJson), &database)

	databaseWriter.EXPECT().Create(database.Name).Return(nil)

	body := strings.NewReader(databaseJson)

	httpRequest, _ := http.NewRequest("POST", "http://localhost/api/v1/databases", body)

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusCreated, httpWriter.Code)
}

func TestPostDatabaseConflict(t *testing.T) {
	container := initResource(t)

	databaseJson := `{"name":"MyDB"}`

	var database *representation.PostDatabase

	json.Unmarshal([]byte(databaseJson), &database)

	databaseWriter.EXPECT().Create(database.Name).Return(errors.New("CONFLICT"))

	body := strings.NewReader(databaseJson)

	httpRequest, _ := http.NewRequest("POST", "http://localhost/api/v1/databases", body)

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"CONFLICT"}`)
}

func TestPostDatabaseUnauthorized(t *testing.T) {
	container := initResource(t)
	databaseJson := `{"name":"MyDB"}`
	body := strings.NewReader(databaseJson)

	httpRequest, _ := http.NewRequest("POST", "http://localhost/api/v1/databases", body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
