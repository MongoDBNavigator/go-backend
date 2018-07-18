package database_resource

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"
	"regexp"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDatabaseSuccess(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")

	databaseWriter.
		EXPECT().
		Delete(dbName).
		Return(nil)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s", dbName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusAccepted, httpWriter.Code)
}

func TestDeleteDatabaseConflict(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")

	databaseWriter.
		EXPECT().
		Delete(dbName).
		Return(errors.New("internal_server_error"))

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s", dbName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}

func TestDeleteDatabaseUnauthorized(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	url := fmt.Sprintf("http://localhost/api/v1/databases/%s", dbName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
