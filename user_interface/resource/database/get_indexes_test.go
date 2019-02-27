package database

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"errors"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetIndexesSuccess(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	var partialFilterExpression interface{}

	indexes := make([]*model.Index, 1)
	indexes[0] = model.NewIndex("_id_", true, true, false, []string{"_id"}, partialFilterExpression)

	indexReader.
		EXPECT().
		ReadAll(dbName, collName).
		Return(indexes, nil)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes", dbName, collName)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusOK, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"objects":[{"name":"_id_","unique":true,"background":true,"sparse":false,"fields":["_id"]}]}`)
}

func TestGetIndexesInternalServerError(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	indexReader.
		EXPECT().
		ReadAll(dbName, collName).
		Return(nil, errors.New("internal_server_error"))

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes", dbName, collName)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}

func TestGetIndexesUnauthorized(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes", dbName, collName)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
