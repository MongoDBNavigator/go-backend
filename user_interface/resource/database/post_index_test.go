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

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/stretchr/testify/assert"
)

func TestPostIndexSuccess(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	indexJson := `{"name":"MyIndex","unique":true,"background":true,"sparse":true,"fields":["name"]}`

	var index *representation.PostIndex

	json.Unmarshal([]byte(indexJson), &index)

	indexModel := model.NewIndex(
		index.Name,
		index.Unique,
		index.Background,
		index.Sparse,
		index.Fields,
	)

	indexWriter.
		EXPECT().
		Create(dbName, collName, indexModel).
		Return(nil)

	body := strings.NewReader(indexJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusCreated, httpWriter.Code)
}

func TestPostIndexConflict(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	indexJson := `{"name":"MyIndex","unique":true,"background":true,"sparse":true,"fields":["name"]}`

	var index *representation.PostIndex

	json.Unmarshal([]byte(indexJson), &index)

	indexModel := model.NewIndex(
		index.Name,
		index.Unique,
		index.Background,
		index.Sparse,
		index.Fields,
	)

	indexWriter.
		EXPECT().
		Create(dbName, collName, indexModel).
		Return(errors.New("CONFLICT"))

	body := strings.NewReader(indexJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"CONFLICT"}`)
}

func TestPostIndexUnauthorized(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	indexJson := `{"name":"MyIndex","unique":true,"background":true,"sparse":true,"fields":["name"]}`

	body := strings.NewReader(indexJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/indexes", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
