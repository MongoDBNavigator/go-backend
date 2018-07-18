package database_resource

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"errors"

	"regexp"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCollectionSuccess(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	collectionsWriter.
		EXPECT().
		Delete(dbName, collName).
		Return(nil)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s", dbName, collName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusAccepted, httpWriter.Code)
}

func TestDeleteCollectionInternalServerError(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	collectionsWriter.
		EXPECT().
		Delete(dbName, collName).
		Return(errors.New("internal_server_error"))

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s", dbName, collName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}

func TestDeleteCollectionUnauthorized(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s", dbName, collName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
