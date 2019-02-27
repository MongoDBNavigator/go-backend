package database

import (
	"testing"

	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"errors"
	"regexp"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestPutDocumentSuccess(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("myDoc")
	docJson := `{"name":"John","gender":"m"}`

	documentReader.
		EXPECT().
		Read(dbName, collName, docId).
		Return(nil, nil)

	documentWriter.
		EXPECT().
		Update(dbName, collName, docId, []byte(docJson)).
		Return(nil)

	body := strings.NewReader(docJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("PUT", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusAccepted, httpWriter.Code)
}

func TestPutDocumentConflict(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("myDoc")

	docJson := `{"name":"John","gender":"m"}`

	documentReader.
		EXPECT().
		Read(dbName, collName, docId).
		Return(nil, nil)

	documentWriter.
		EXPECT().
		Update(dbName, collName, docId, []byte(docJson)).
		Return(errors.New("CONFLICT"))

	body := strings.NewReader(docJson)

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("PUT", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusConflict, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"CONFLICT"}`)
}

func TestPutDocumentNotFound(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("myDoc")

	docJson := `{"name":"John","gender":"m"}`

	documentReader.
		EXPECT().
		Read(dbName, collName, docId).
		Return(nil, errors.New("NOT_FOUND"))

	body := strings.NewReader(docJson)
	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("PUT", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusNotFound, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"NOT_FOUND"}`)
}

func TestPutDocumentUnauthorized(t *testing.T) {
	container := initResource(t)

	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")
	docId := value.DocId("myDoc")

	docJson := `{"name":"John","gender":"m"}`

	body := strings.NewReader(docJson)
	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/documents/%s", dbName, collName, docId)

	httpRequest, _ := http.NewRequest("PUT", url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.ServeHTTP(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
