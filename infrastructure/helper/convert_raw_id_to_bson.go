package helper

import (
	"log"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Convert string document ID representation to bson.Element
// https://docs.mongodb.com/manual/reference/method/ObjectId/
func ConvertStringIDToBJSON(docID string) *bson.D {
	id, err := primitive.ObjectIDFromHex(docID)
	var element *bson.D

	if err != nil {
		log.Println(err)
		if i, err := strconv.Atoi(docID); err == nil {
			element = &bson.D{{"_id", i}}
		} else {
			log.Println(err)
			element = &bson.D{{"_id", docID}}
		}
	} else {
		element = &bson.D{{"_id", id}}
	}

	return element
}
