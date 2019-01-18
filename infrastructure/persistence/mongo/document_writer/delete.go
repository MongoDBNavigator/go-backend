package document_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/infrastructure/helper"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Drop document
// https://docs.mongodb.com/manual/tutorial/remove-documents/
func (rcv *documentWriter) Delete(dbName value.DBName, collName value.CollName, docId value.DocId) error {
	_, err := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		DeleteOne(context.Background(), helper.ConvertStringIDToBJSON(string(docId)))

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
