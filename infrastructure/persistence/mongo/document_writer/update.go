package document_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/infrastructure/helper"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Update document
// https://docs.mongodb.com/manual/tutorial/update-documents/
func (rcv *documentWriter) Update(dbName value.DBName, collName value.CollName, docId value.DocId, doc []byte) error {
	document := bson.D{}

	if err := bson.UnmarshalExtJSON(doc, true, &document); err != nil {
		return err
	}

	_, err := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		ReplaceOne(
			context.Background(),
			helper.ConvertStringIDToBJSON(string(docId)),
			document,
		)

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
