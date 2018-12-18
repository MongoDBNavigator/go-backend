package collection_reader

import (
	"context"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Read all collections from DB
func (rcv *collectionReader) ReadAll(dbName value.DBName) ([]*model.Collection, error) {
	cursor, err := rcv.db.Database(string(dbName)).ListCollections(context.Background(), bson.NewDocument())

	if err != nil {
		return nil, err
	}

	result := make([]*model.Collection, 0)

	collection := bson.NewDocument()
	for cursor.Next(context.Background()) {
		collection.Reset() // Reset clears a document so it can be reused
		if err := cursor.Decode(collection); err != nil {
			return nil, err
		}

		name := collection.LookupElement("name")
		// The collStats command returns a variety of storage statistics for a given collection.
		// https://docs.mongodb.com/manual/reference/command/collStats/
		collStats, err := rcv.db.Database(string(dbName)).RunCommand(
			context.Background(),
			bson.NewDocument(bson.EC.String("collStats", name.Value().StringValue())),
		)

		if err != nil {
			return nil, err
		}

		var count, avgObjSize, indexesNumber, size int

		// The total uncompressed size in memory of all records in a collection
		if countRaw, _ := collStats.Lookup("count"); countRaw != nil {
			count = int(countRaw.Value().Int32())
		}

		// The number of indexes on the collection
		if nIndexes, _ := collStats.Lookup("nindexes"); nIndexes != nil {
			indexesNumber = int(nIndexes.Value().Int32())
		}

		// The total uncompressed size in memory of all records in a collection
		if sizeRaw, _ := collStats.Lookup("size"); sizeRaw != nil {
			size = int(sizeRaw.Value().Int32())
		}

		// The average size of an object in the collection
		if avgObjSizeRaw, _ := collStats.Lookup("avgObjSize"); avgObjSizeRaw != nil {
			avgObjSize = int(avgObjSizeRaw.Value().Int32())
		}

		result = append(result, model.NewCollection(name.Value().StringValue(), count, indexesNumber, avgObjSize, size))
	}

	return result, nil
}
