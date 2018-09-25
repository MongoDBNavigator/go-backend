package database_reader

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
)

// Returns a document that lists all databases and returns basic database statistics.
// https://docs.mongodb.com/manual/reference/command/listDatabases/
func (rcv *databaseReader) ReadAll() ([]*model.Database, error) {
	databaseNames, err := rcv.db.ListDatabaseNames(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Database, len(databaseNames))

	for i, name := range databaseNames {
		dbStats, err := rcv.db.Database(name).RunCommand(
			context.Background(),
			bson.NewDocument(bson.EC.Int32("dbStats", 1)),
		)

		if err != nil {
			return nil, err
		}

		var collections, indexesNumber, storageSize int

		if collectionsRaw, _ := dbStats.Lookup("collections"); collectionsRaw != nil {
			collections = int(collectionsRaw.Value().Int32())
		}

		if indexes, _ := dbStats.Lookup("indexes"); indexes != nil {
			indexesNumber = int(indexes.Value().Int32())
		}

		if storageSizeRaw, _ := dbStats.Lookup("storageSize"); storageSizeRaw != nil {
			storageSize = int(storageSizeRaw.Value().Double())
		}

		result[i] = model.NewDatabase(name, storageSize, indexesNumber, collections)
	}

	return result, nil
}
