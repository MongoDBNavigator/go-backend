package collection_reader

import (
	"context"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Read all collections from DB
func (rcv *collectionReader) ReadAll(dbName value.DBName) ([]*model.Collection, error) {
	cursor, err := rcv.db.Database(string(dbName)).ListCollections(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	result := make([]*model.Collection, 0)

	for cursor.Next(context.Background()) {
		collection := new(collectionStructure)
		if err := cursor.Decode(&collection); err != nil {
			return nil, err
		}

		// The collStats command returns a variety of storage statistics for a given collection.
		// https://docs.mongodb.com/manual/reference/command/collStats/
		collStatsResult := rcv.db.Database(string(dbName)).RunCommand(
			context.Background(),
			bson.D{{"collStats", collection.Name}},
		)

		if collStatsResult.Err() != nil {
			return nil, collStatsResult.Err()
		}

		raw, err := collStatsResult.DecodeBytes()

		if err != nil {
			return nil, err
		}

		var count, avgObjSize, indexesNumber, size int

		// The total uncompressed size in memory of all records in a collection
		if countRaw, err := raw.LookupErr("count"); err == nil {
			count = int(countRaw.Int32())
		}

		// The number of indexes on the collection
		if nIndexes, err := raw.LookupErr("nindexes"); err == nil {
			indexesNumber = int(nIndexes.Int32())
		}

		// The total uncompressed size in memory of all records in a collection
		if sizeRaw, err := raw.LookupErr("size"); err == nil {
			size = int(sizeRaw.Int32())
		}

		// The average size of an object in the collection
		if avgObjSizeRaw, err := raw.LookupErr("avgObjSize"); err == nil {
			avgObjSize = int(avgObjSizeRaw.Int32())
		}

		result = append(result, model.NewCollection(collection.Name, count, indexesNumber, avgObjSize, size))
	}

	return result, nil
}
