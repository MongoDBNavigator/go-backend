package document_writer

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Drop document
// https://docs.mongodb.com/manual/tutorial/remove-documents/
func (rcv *documentWriter) Delete(dbName value.DBName, collName value.CollName, docId value.DocId) error {
	id, err := objectid.FromHex(string(docId))

	var element *bson.Element

	if err != nil {
		log.Println(err)
		element = bson.EC.String("_id", string(docId))
	} else {
		element = bson.EC.ObjectID("_id", id)
	}

	_, err = rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		DeleteOne(context.Background(), bson.NewDocument(element))

	if err != nil {
		return err
	}

	return nil
}
