package database_resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"errors"
	"regexp"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestPostDocumentSuccess(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	docJson := `{"name":"John","gender":"m"}`

	var doc interface{}

	json.Unmarshal([]byte(docJson), &doc)

	documentWriter.
		EXPECT().
		Create(dbName, collName, &doc).
		Return(nil)

	body := strings.NewReader(docJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusCreated, httpWriter.Code)
}

func TestPostDocumentConflict(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	docJson := `{"name":"John","gender":"m"}`

	var doc interface{}

	json.Unmarshal([]byte(docJson), &doc)

	documentWriter.
		EXPECT().
		Create(dbName, collName, &doc).
		Return(errors.New("CONFLICT"))

	body := strings.NewReader(docJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"CONFLICT"}`)
}

func TestPostDocumentUnauthorized(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	docJson := `{"name":"John","gender":"m"}`

	body := strings.NewReader(docJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
