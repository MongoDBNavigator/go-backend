package index_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Returns an array of documents that describe the existing indexes on a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.getIndexes/#db.collection.getIndexes
func (rcv *indexReader) ReadAll(dbName value.DBName, collName value.CollName) ([]*model.Index, error) {
	indexes, err := rcv.db.DB(string(dbName)).C(string(collName)).Indexes()

	if err != nil {
		return nil, err
	}

	var partialFilterExpression interface{}

	result := make([]*model.Index, len(indexes))

	for i, index := range indexes {
		result[i] = model.NewIndex(
			index.Name,
			index.Unique,
			index.Background,
			index.Sparse,
			index.Key,
			partialFilterExpression,
		)
	}

	return result, nil
}
