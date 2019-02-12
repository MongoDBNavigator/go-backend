package request

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Extract collection name from url path
func ExtractCollectionName(r *http.Request) (value.CollName, error) {
	vars := mux.Vars(r)
	collName := value.CollName(vars["collectionName"])

	if err := collName.Valid(); err != nil {
		return "", err
	}

	return collName, nil
}

// Extract database name from url path
func ExtractDatabaseName(r *http.Request) (value.DBName, error) {
	vars := mux.Vars(r)
	dbName := value.DBName(vars["databaseName"])

	if err := dbName.Valid(); err != nil {
		return "", err
	}

	return dbName, nil
}

// Extract documentId from url path
func ExtractDocumentId(r *http.Request) (value.DocId, error) {
	vars := mux.Vars(r)
	docId := value.DocId(vars["documentId"])

	if err := docId.Valid(); err != nil {
		return "", err
	}

	return docId, nil
}

// Extract index name from url path
func ExtractIndex(r *http.Request) (value.IndexName, error) {
	vars := mux.Vars(r)
	indexName := value.IndexName(vars["indexName"])

	if err := indexName.Valid(); err != nil {
		return "", err
	}

	return indexName, nil
}

// Extract all allowed parameters from url path
func ExtractParametersFromRequest(request *http.Request, db *value.DBName, coll *value.CollName, docId *value.DocId, index *value.IndexName) error {
	if db != nil {
		dbName, err := ExtractDatabaseName(request)

		if err != nil {
			return err
		}

		*db = dbName
	}

	if coll != nil {
		collName, err := ExtractCollectionName(request)

		if err != nil {
			return err
		}

		*coll = collName
	}

	if docId != nil {
		id, err := ExtractDocumentId(request)

		if err != nil {
			return err
		}

		*docId = id
	}

	if index != nil {
		indexName, err := ExtractIndex(request)

		if err != nil {
			return err
		}

		*index = indexName
	}

	return nil
}
