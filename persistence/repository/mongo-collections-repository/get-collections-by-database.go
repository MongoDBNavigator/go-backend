package mongo_collections_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
)

// Retrieve information, i.e. the name and options, about the collections and views in a database.
// https://docs.mongodb.com/manual/reference/command/listCollections/
func (rcv *collectionsRepository) GetCollectionsByDatabase(databaseName string) ([]*model.Collection, error) {
	collectionNames, err := rcv.db.DB(databaseName).CollectionNames()

	if err != nil {
		return nil, err
	}

	result := make([]*model.Collection, len(collectionNames))

	for i, name := range collectionNames {

		stats, err := rcv.GetStats(databaseName, name)

		if err != nil {
			return nil, err
		}

		result[i] = model.NewCollection(name, stats.GetCount(), stats.GetIndexesNumber(), stats.GetAvgObjSize())
	}

	return result, nil
}
