package document_reader

import (
	"context"
	"encoding/json"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Fetch document by ID
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) Read(dbName value.DBName, collName value.CollName, docId value.DocId) (interface{}, error) {
	id, err := objectid.FromHex(string(docId))

	var element *bson.Element

	if err != nil {
		log.Println(err)
		element = bson.EC.String("_id", string(docId))
	} else {
		element = bson.EC.ObjectID("_id", id)
	}

	document := bson.NewDocument()

	err = rcv.db.
		Database(string(dbName)).
		Collection(string(collName)).
		FindOne(context.Background(), bson.NewDocument(element)).
		Decode(document)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	docJson, err := document.ToExtJSONErr(false)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result map[string]interface{}

	if err := json.Unmarshal([]byte(docJson), &result); err != nil {
		log.Println(err)
		return nil, err
	}

	if objID, ok := result["_id"]; ok {
		result["_id"] = rcv.objectIDToScalarType(objID)
	}

	return result, nil
}
