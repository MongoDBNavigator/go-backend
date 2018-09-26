package document_writer

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Insert new document
// https://docs.mongodb.com/manual/tutorial/insert-documents/
func (rcv *documentWriter) Create(dbName value.DBName, collName value.CollName, doc interface{}) error {
	document, err := bson.NewDocumentEncoder().EncodeDocument(doc)

	fmt.Println(err)
	fmt.Println(document)

	_, err = rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		InsertOne(context.Background(), document)

	fmt.Println(err)

	return err
}
