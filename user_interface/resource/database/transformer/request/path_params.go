package request

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/emicklei/go-restful"
)

// Extract collection name from url path
func ExtractCollectionName(request *restful.Request) (value.CollName, error) {
	collName := value.CollName(request.PathParameter("collectionName"))

	if err := collName.Valid(); err != nil {
		return "", err
	}

	return collName, nil
}

// Extract database name from url path
func ExtractDatabaseName(request *restful.Request) (value.DBName, error) {
	dbName := value.DBName(request.PathParameter("databaseName"))

	if err := dbName.Valid(); err != nil {
		return "", err
	}

	return dbName, nil
}

// Extract documentId from url path
func ExtractDocumentId(request *restful.Request) (value.DocId, error) {
	docId := value.DocId(request.PathParameter("documentId"))

	if err := docId.Valid(); err != nil {
		return "", err
	}

	return docId, nil
}

// Extract index name from url path
func ExtractIndex(request *restful.Request) (value.IndexName, error) {
	indexName := value.IndexName(request.PathParameter("indexName"))

	if err := indexName.Valid(); err != nil {
		return "", err
	}

	return indexName, nil
}

// Extract all allowed parameters from url path
func ExtractParametersFromRequest(request *restful.Request, db *value.DBName, coll *value.CollName, docId *value.DocId, index *value.IndexName) error {
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
