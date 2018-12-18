package document_writer

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Update document
// https://docs.mongodb.com/manual/tutorial/update-documents/
func (rcv *documentWriter) Update(dbName value.DBName, collName value.CollName, docId value.DocId, doc interface{}) error {
	document, err := bson.NewDocumentEncoder().EncodeDocument(doc)

	if err != nil {
		log.Println(err)
		return err
	}

	id, err := objectid.FromHex(string(docId))

	var element *bson.Element

	if err != nil {
		log.Println(err)
		element = bson.EC.Interface("_id", docId)
	} else {
		element = bson.EC.ObjectID("_id", id)
	}

	_, err = rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		ReplaceOne(context.Background(), bson.NewDocument(element), document)

	if err != nil {
		log.Println(err)
	}

	return err
}
