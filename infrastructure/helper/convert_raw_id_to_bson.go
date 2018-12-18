package helper

import (
	"log"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Convert string document ID representation to bson.Element
func ConvertStringIDToBJSON(docID string) *bson.Element {
	id, err := objectid.FromHex(docID)
	var element *bson.Element

	if err != nil {
		log.Println(err)
		if i, err := strconv.Atoi(docID); err == nil {
			element = bson.EC.Int64("_id", int64(i))
		} else {
			log.Println(err)
			element = bson.EC.String("_id", docID)
		}
	} else {
		element = bson.EC.ObjectID("_id", id)
	}

	return element
}
