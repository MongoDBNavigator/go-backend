package collection_reader

import (
	"context"
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

func (rcv *collectionReader) ReadAll(dbName value.DBName) ([]*model.Collection, error) {
	cursor, err := rcv.db.Database(string(dbName)).ListCollections(context.Background(), bson.NewDocument())

	if err != nil {
		return nil, err
	}

	result := make([]*model.Collection, 0)

	collection := bson.NewDocument()
	for cursor.Next(context.Background()) {
		collection.Reset()
		if err := cursor.Decode(collection); err != nil {
			return nil, err
		}

		name := collection.LookupElement("name")
		// The collStats command returns a variety of storage statistics for a given collection.
		// https://docs.mongodb.com/manual/reference/command/collStats/
		collStats, err := rcv.db.Database(string(dbName)).RunCommand(
			context.Background(),
			bson.NewDocument(
				bson.EC.String("collStats", name.Value().StringValue()),
			),
		)

		if err != nil {
			return  nil, err
		}

		var count, avgObjSize, indexesNumber, size int

		if countRaw, err := collStats.Lookup("count"); err != nil {
			count = int(countRaw.Value().Int32())
		}

		if nIndexes, err := collStats.Lookup("nindexes"); err != nil {
			indexesNumber = int(nIndexes.Value().Int32())
		}

		if sizeRaw, err := collStats.Lookup("size"); err != nil {
			size = int(sizeRaw.Value().Int32())
		}

		if avgObjSizeRaw, err := collStats.Lookup("avgObjSize"); err != nil && avgObjSizeRaw != nil {
			avgObjSize = int(avgObjSizeRaw.Value().Int32())
		}

		result = append(result, model.NewCollection(name.Value().StringValue(), count, indexesNumber, avgObjSize, size))
	}

	return result, nil
}
