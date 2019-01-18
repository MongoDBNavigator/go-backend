package document_reader

import (
	"context"
	"encoding/json"
	"log"

	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Fetch documents with pagination and filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadAll(conditions *value.ReadAllDocConditions) ([]interface{}, error) {
	opts := options.Find()
	opts.SetSkip(int64(conditions.Skip()))
	opts.SetLimit(int64(conditions.Skip()))

	if len(conditions.Sort()) > 0 {
		opts.SetSort(conditions.Sort())
	}

	var filterDoc bsonx.Doc

	if len(conditions.Filter()) > 0 {
		filterDoc = rcv.convertFilterToBson(conditions.Filter())
	}

	cursor, err := rcv.db.
		Database(string(conditions.DbName())).
		Collection(string(conditions.CollName())).
		Find(context.Background(), filterDoc, opts)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []interface{}

	for cursor.Next(context.Background()) {
		var document bsonx.Doc

		if err := cursor.Decode(&document); err != nil {
			log.Println(err)
			return nil, err
		}

		raw, err := bson.MarshalExtJSON(document, false, false)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		var doc map[string]interface{}

		if err := json.Unmarshal(raw, &doc); err != nil {
			log.Println(err)
			return nil, err
		}

		if objID, ok := doc["_id"]; ok {
			doc["_id"] = rcv.objectIDToScalarType(objID)
		}

		result = append(result, doc)
	}

	return result, nil
}
