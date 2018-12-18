package database

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"
	"fmt"
	"regexp"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetDatabasesSuccess(t *testing.T) {
	container := initResource(t)

	dbs := make([]*model.Database, 1)
	dbs[0] = model.NewDatabase("MyDB", 1, 2, 3)

	databaseReader.
		EXPECT().
		ReadAll().
		Return(dbs, nil)

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/databases", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusOK, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"objects":[{"name":"MyDB","collections":3,"indexesNumber":2,"storageSize":1}]}`)
}

func TestGetDatabasesInternalServerError(t *testing.T) {
	container := initResource(t)

	databaseReader.
		EXPECT().
		ReadAll().
		Return(nil, errors.New("internal_server_error"))

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/databases", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}

func TestGetDatabasesUnauthorized(t *testing.T) {
	container := initResource(t)

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/databases", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
