package mongo_collections_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
)

//
// Returns an array of documents that describe the existing indexes on a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.getIndexes/#db.collection.getIndexes
//
func (rcv *collectionsRepository) GetIndexes(databaseName string, collectionName string) ([]*model.Index, error) {
	indexes, err := rcv.db.DB(databaseName).C(collectionName).Indexes()

	if err != nil {
		return nil, err
	}

	result := make([]*model.Index, len(indexes))

	for i, index := range indexes {
		result[i] = model.NewIndex(
			index.Name,
			index.Unique,
			index.Background,
			index.Sparse,
			index.Key,
		)
	}

	return result, nil
}
