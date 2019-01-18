package index_reader

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/x/bsonx"

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

	for cursor.Next(context.Background()) {
		var index bsonx.Doc

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

		if val, err := index.LookupErr("name"); err == nil {
			name = val.StringValue()
		}

		if val, err := index.LookupErr("background"); err == nil {
			background = val.Boolean()
		}

		if val, err := index.LookupErr("unique"); err == nil {
			unique = val.Boolean()
		}

		if val, err := index.LookupErr("sparse"); err == nil {
			sparse = val.Boolean()
		}

		rawKeys, err := index.LookupErr("key")

		if err != nil {
			log.Println(err)
			return nil, err
		}

		keys := make([]string, len(rawKeys.Document()))

		for i, key := range rawKeys.Document() {
			keys[i] = key.Key
		}

		result = append(result, model.NewIndex(name, unique, background, sparse, keys, partialFilterExpression))
	}

	return result, nil
}
