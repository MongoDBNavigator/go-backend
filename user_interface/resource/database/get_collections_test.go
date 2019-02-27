package database

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"regexp"

	"errors"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetCollectionsSuccess(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")

	colls := make([]*model.Collection, 1)
	colls[0] = model.NewCollection("MyColl", 1, 1, 1, 1)

	collectionsReader.
		EXPECT().
		ReadAll(dbName).
		Return(colls, nil)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections", dbName)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusOK, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(
		t,
		`{"objects":[{"name":"MyColl","size":1,"documentsNumber":1,"indexesNumber":1,"avgObjSize":1}]}`,
		space.ReplaceAllString(httpWriter.Body.String(), ""),
	)

}

func TestGetCollectionsInternalServerError(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")

	collectionsReader.
		EXPECT().
		ReadAll(dbName).
		Return(nil, errors.New("internal_server_error"))

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections", dbName)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}

func TestGetCollectionsUnauthorized(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections", dbName)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
