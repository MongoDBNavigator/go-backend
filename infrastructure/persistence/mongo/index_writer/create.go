package index_writer

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Creates index on collections.
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
func (rcv *indexWriter) Create(dbName value.DBName, collName value.CollName, index *model.Index) error {
	keys := bson.NewDocument()

	for _, indexName := range index.Fields() {
		keys.Set(bson.EC.Int32(indexName, 1))
	}

	result, err := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		Indexes().
		CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys: keys,
				Options: mongo.
					NewIndexOptionsBuilder().
					Name(index.Name()).
					Unique(index.Unique()).
					Background(index.Background()).
					Sparse(index.Sparse()).
					Build(),
			},
		)

	log.Println(result)

	if err != nil {
		log.Println(err)
	}

	return err
}
