package transformer

import (
	"github.com/emicklei/go-restful"
)

func ExtractParametersFromRequest(request *restful.Request, db *string, coll *string, docId *string, index *string) error {
	if db != nil {
		dbName, err := RequestToDatabaseName(request)

		if err != nil {
			return err
		}

		*db = dbName
	}

	if coll != nil {
		collName, err := RequestToCollectionName(request)

		if err != nil {
			return err
		}

		*coll = collName
	}

	if docId != nil {
		id, err := RequestToDocumentId(request)

		if err != nil {
			return err
		}

		*docId = id
	}

	if index != nil {
		indexName, err := RequestToIndex(request)

		if err != nil {
			return err
		}

		*index = indexName
	}

	return nil
}
