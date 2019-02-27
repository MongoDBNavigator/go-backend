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

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/stretchr/testify/assert"
)

func TestPostCollectionSuccess(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	indexJson := `{"name":"MyCollection"}`

	var coll *representation.PostCollection

	json.Unmarshal([]byte(indexJson), &coll)

	collectionsWriter.
		EXPECT().
		Create(dbName, coll.Name).
		Return(nil)

	body := strings.NewReader(indexJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections", dbName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusCreated, httpWriter.Code)
}

func TestPostCollectionConflict(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	indexJson := `{"name":"MyCollection"}`

	var coll *representation.PostCollection

	json.Unmarshal([]byte(indexJson), &coll)

	collectionsWriter.
		EXPECT().
		Create(dbName, coll.Name).
		Return(errors.New("CONFLICT"))

	body := strings.NewReader(indexJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections", dbName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"CONFLICT"}`)
}

func TestPostCollectionUnauthorized(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	indexJson := `{"name":"MyCollection"}`
	body := strings.NewReader(indexJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections", dbName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
