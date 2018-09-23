package index_reader

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

type indexSpec struct {
	Name       string        "name,omitempty"
	Unique     bool          "unique,omitempty"
	Background bool          "background,omitempty"
	Sparse     bool          "sparse,omitempty"
	Key        bson.Document "key,omitempty"
}

// Returns an array of documents that describe the existing indexes on a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.getIndexes/#db.collection.getIndexes
func (rcv *indexReader) ReadAll(dbName value.DBName, collName value.CollName) ([]*model.Index, error) {
	indexesCursor, err := rcv.db.Database(string(dbName)).Collection(string(collName)).Indexes().List(context.Background())

	if err != nil {
		return nil, err
	}

	var partialFilterExpression interface{}
	result := make([]*model.Index, 0)

	for indexesCursor.Next(context.Background()) {
		index := indexSpec{}

		if err := indexesCursor.Decode(&index); err != nil {
			fmt.Println(err)
			return nil, err
		}

		keys := make([]string, 0)

		var i uint
		for {
			key, ok := index.Key.ElementAtOK(i)
			if !ok {
				break
			}

			keys = append(keys, key.Key())
			i++
		}

		result = append(result, model.NewIndex(
			index.Name,
			index.Unique,
			index.Background,
			index.Sparse,
			keys,
			partialFilterExpression,
		))
	}

	return result, nil
}
