package document_writer

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Drop document
// https://docs.mongodb.com/manual/tutorial/remove-documents/
func (rcv *documentWriter) Delete(dbName value.DBName, collName value.CollName, docId value.DocId) error {
	res := rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		FindOneAndDelete(context.Background(), bson.NewDocument(
			bson.EC.String("_id", string(docId)),
		))

	//fmt.Println(err)
	fmt.Println(res)

	return nil
}
