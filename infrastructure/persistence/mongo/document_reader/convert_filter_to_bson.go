package document_reader

import (
	"log"

	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/mongodb/mongo-go-driver/bson"
)

// https://docs.mongodb.com/manual/reference/operator/query-comparison/
func (rcv *documentReader) convertFilterToBson(data []byte) bsonx.Doc {
	document := bsonx.Doc{}

	if err := bson.UnmarshalExtJSON(data, false, &document); err != nil {
		log.Println(err)

		return document
	}

	for i, element := range document {
		log.Println(i, element)
	}

	return document
}
