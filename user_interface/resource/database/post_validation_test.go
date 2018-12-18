package database

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/stretchr/testify/assert"
)

func TestPostValidationUnauthorized(t *testing.T) {
	container := initResource(t)
	dbName := value.DBName("myDB")
	collName := value.CollName("myColl")

	url := fmt.Sprintf("http://localhost/api/v1/databases/%s/collections/%s/validation", dbName, collName)

	httpRequest, _ := http.NewRequest("POST", url, nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	container.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}
