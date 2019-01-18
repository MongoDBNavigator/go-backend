package collection_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Creates a new collection or view.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *collectionWriter) Create(dbName value.DBName, collName value.CollName) error {
	result := rcv.db.Database(string(dbName)).RunCommand(
		context.Background(),
		bson.D{{"create", string(collName)}},
	)

	if result.Err() != nil {
		log.Println(result.Err())
	}

	return result.Err()
}
