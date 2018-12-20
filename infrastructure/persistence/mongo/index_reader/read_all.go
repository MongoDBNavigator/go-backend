package index_reader

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Returns an array of documents that describe the existing indexes on a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.getIndexes/#db.collection.getIndexes
func (rcv *indexReader) ReadAll(dbName value.DBName, collName value.CollName) ([]*model.Index, error) {
	cursor, err := rcv.db.Database(string(dbName)).Collection(string(collName)).Indexes().List(context.Background())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var partialFilterExpression interface{}
	result := make([]*model.Index, 0)
	index := bson.NewDocument()
	for cursor.Next(context.Background()) {
		index.Reset()

		if err := cursor.Decode(&index); err != nil {
			log.Println(err)
			return nil, err
		}

		var (
			name       string
			unique     bool
			background bool
			sparse     bool
		)

		if val := index.Lookup("name"); val != nil {
			name = val.StringValue()
		}

		if val := index.Lookup("background"); val != nil {
			background = val.Boolean()
		}

		if val := index.Lookup("unique"); val != nil {
			unique = val.Boolean()
		}

		if val := index.Lookup("sparse"); val != nil {
			sparse = val.Boolean()
		}

		rawKeys, err := index.Lookup("key").MutableDocument().Keys(false)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		keys := make([]string, len(rawKeys))

		for i, key := range rawKeys {
			keys[i] = key.Name
		}

		result = append(result, model.NewIndex(
			name,
			unique,
			background,
			sparse,
			keys,
			partialFilterExpression,
		))
	}

	return result, nil
}
