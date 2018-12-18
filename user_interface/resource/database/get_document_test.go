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

func TestGetDocumentSuccess(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("id")

	doc := struct {
		Name string `json:"name"`
	}{Name: "John"}

	documentReader.
		EXPECT().
		Read(dbName, collName, docId).
		Return(doc, nil)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusOK, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"name":"John"}`)
}

func TestGetDocumentNotFound(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("id")

	documentReader.
		EXPECT().
		Read(dbName, collName, docId).
		Return(nil, errors.New("NOT_FOUND"))

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusNotFound, httpWriter.Code)
}

func TestGetDocumentUnauthorized(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("id")

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
