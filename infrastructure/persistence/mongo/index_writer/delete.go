package index_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Drops or removes the specified index from a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.dropIndex/#db.collection.dropIndex
func (rcv *indexWriter) Delete(dbName value.DBName, collName value.CollName, indexName value.IndexName) error {
	_, err := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		Indexes().
		DropOne(context.Background(), string(indexName))

	if err != nil {
		log.Println(err)
	}

	return err
}
