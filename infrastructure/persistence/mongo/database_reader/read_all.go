package database_reader

import (
	"context"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Returns a document that lists all databases and returns basic database statistics.
// https://docs.mongodb.com/manual/reference/command/listDatabases/
func (rcv *databaseReader) ReadAll() ([]*model.Database, error) {
	databaseNames, err := rcv.db.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	result := make([]*model.Database, len(databaseNames))

	for i, name := range databaseNames {
		dbStatsResult := rcv.db.Database(name).RunCommand(
			context.Background(),
			bson.D{{"dbStats", 1}},
		)

		if dbStatsResult.Err() != nil {
			return nil, dbStatsResult.Err()
		}

		raw, err := dbStatsResult.DecodeBytes()

		if err != nil {
			return nil, err
		}

		var collections, indexesNumber, storageSize int

		// Contains a count of the number of collections in that database.
		if collectionsRaw, err := raw.LookupErr("collections"); err == nil {
			collections = int(collectionsRaw.Int32())
		}

		// Contains a count of the total number of indexes across all collections in the database.
		if indexes, err := raw.LookupErr("indexes"); err == nil {
			indexesNumber = int(indexes.Int32())
		}

		// The total amount of space allocated to collections in this database for document storage.
		if storageSizeRaw, err := raw.LookupErr("storageSize"); err == nil {
			storageSize = int(storageSizeRaw.Double())
		}

		result[i] = model.NewDatabase(name, storageSize, indexesNumber, collections)
	}

	return result, nil
}
