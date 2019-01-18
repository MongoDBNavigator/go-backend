package document_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Insert new document
// https://docs.mongodb.com/manual/tutorial/insert-documents/
func (rcv *documentWriter) Create(dbName value.DBName, collName value.CollName, doc []byte) error {
	document := bson.D{}

	if err := bson.UnmarshalExtJSON(doc, false, &document); err != nil {
		log.Println(err)
		return err
	}

	_, err := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		InsertOne(context.Background(), document)

	if err != nil {
		log.Println(err)
	}

	return err
}
