package document_reader

import (
	"context"
	"encoding/json"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Fetch documents with pagination and filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadAll(conditions *value.ReadAllDocConditions) ([]interface{}, error) {
	opts := make([]findopt.Find, 2)
	opts[0] = findopt.Skip(int64(conditions.Skip()))
	opts[1] = findopt.Limit(int64(conditions.Limit()))

	if len(conditions.Sort()) > 0 {
		opts = append(opts, findopt.Sort(conditions.Sort()))
	}

	cursor, err := rcv.db.
		Database(string(conditions.DbName())).
		Collection(string(conditions.CollName())).
		Find(context.Background(), nil, opts...)

	if err != nil {
		return nil, err
	}

	var result []interface{}
	document := bson.NewDocument()

	for cursor.Next(nil) {
		document.Reset()
		if err := cursor.Decode(document); err != nil {
			log.Println(err)
			return nil, err
		}

		docJson, err := document.ToExtJSONErr(false)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		var doc map[string]interface{}

		if err := json.Unmarshal([]byte(docJson), &doc); err != nil {
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
