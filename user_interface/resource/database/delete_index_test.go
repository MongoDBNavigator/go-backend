package database

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

func TestDeleteIndexSuccess(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	indexName := value.IndexName("myDoc")

	indexWriter.
		EXPECT().
		Delete(dbName, collName, indexName).
		Return(nil)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes/%s", dbName, collName, indexName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusAccepted, httpWriter.Code)
}

func TestDeleteIndexConflict(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	indexName := value.IndexName("myDoc")

	indexWriter.
		EXPECT().
		Delete(dbName, collName, indexName).
		Return(errors.New("internal_server_error"))

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes/%s", dbName, collName, indexName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}

func TestDeleteIndexUnauthorized(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	indexName := value.IndexName("myDoc")

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes/%s", dbName, collName, indexName)

	httpRequest, _ := http.NewRequest("DELETE", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
