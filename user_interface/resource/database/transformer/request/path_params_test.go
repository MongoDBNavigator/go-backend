package request

import (
	"net/http"
	"testing"

	"github.com/emicklei/go-restful"
	"github.com/stretchr/testify/assert"
)

func TestExtractCollectionNameSuccess(t *testing.T) {
	collectionName := "test_collection_name"
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["collectionName"] = collectionName

	value, err := ExtractCollectionName(req)

	assert.Nil(t, err)
	assert.EqualValues(t, collectionName, value)
}

func TestExtractCollectionNameFail(t *testing.T) {
	collectionName := ""
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["collectionName"] = collectionName

	_, err := ExtractCollectionName(req)

	assert.Error(t, err)
}

func TestExtractDatabaseNameSuccess(t *testing.T) {
	dbName := "test_db_name"
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["databaseName"] = dbName

	value, err := ExtractDatabaseName(req)

	assert.Nil(t, err)
	assert.EqualValues(t, dbName, value)
}

func TestExtractDatabaseNameFail(t *testing.T) {
	dbName := ""
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["databaseName"] = dbName

	_, err := ExtractDatabaseName(req)

	assert.Error(t, err)
}

func TestExtractDocumentIdSuccess(t *testing.T) {
	dbName := "test_doc_id"
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["documentId"] = dbName

	value, err := ExtractDocumentId(req)

	assert.Nil(t, err)
	assert.EqualValues(t, dbName, value)
}

func TestExtractDocumentIdFail(t *testing.T) {
	dbName := ""
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["documentId"] = dbName

	_, err := ExtractDocumentId(req)

	assert.Error(t, err)
}

func TestExtractIndexSuccess(t *testing.T) {
	indexName := "test_index_name"
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["indexName"] = indexName

	value, err := ExtractIndex(req)

	assert.Nil(t, err)
	assert.EqualValues(t, indexName, value)
}

func TestExtractIndexFail(t *testing.T) {
	dbName := ""
	req := restful.NewRequest(&http.Request{Method: "GET"})
	req.PathParameters()["indexName"] = dbName

	_, err := ExtractIndex(req)

	assert.Error(t, err)
}
