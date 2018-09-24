package collection_writer

import (
	"context"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Creates a new collection or view.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *collectionWriter) Create(dbName value.DBName, collName value.CollName) error {
	_, err := rcv.db.Database(string(dbName)).RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.String("create", string(collName))),
	)

	return err
}
