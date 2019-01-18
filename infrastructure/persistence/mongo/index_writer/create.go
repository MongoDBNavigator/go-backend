package index_writer

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Creates index on collections.
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
func (rcv *indexWriter) Create(dbName value.DBName, collName value.CollName, index *model.Index) error {
	keys := bsonx.Doc{}

	for _, indexName := range index.Fields() {
		keys = keys.Append(indexName, bsonx.Int32(1))
	}

	_, err := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		Indexes().
		CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys: &keys,
				Options: options.Index().
					SetName(index.Name()).
					SetUnique(index.Unique()).
					SetBackground(index.Background()).
					SetSparse(index.Sparse()),
			},
		)

	if err != nil {
		log.Println(err)
	}

	return err
}
