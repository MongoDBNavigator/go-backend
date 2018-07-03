package mongo_collections_repository

import (
	"fmt"

	"strconv"

	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"gopkg.in/mgo.v2/bson"
)

// The collStats command returns a variety of storage statistics for a given collection.
// https://docs.mongodb.com/manual/reference/command/collStats/
func (rcv *collectionsRepository) GetStats(databaseName string, collectionName string) (*model.CollectionStats, error) {
	var stats bson.M

	if err := rcv.db.DB(databaseName).Run(bson.D{{"collStats", &collectionName}}, &stats); err != nil {
		return nil, err
	}

	ns := fmt.Sprintf("%v", stats["ns"])
	sizeString := fmt.Sprintf("%v", stats["size"])
	countString := fmt.Sprintf("%v", stats["count"])
	cappedString := fmt.Sprintf("%v", stats["capped"])
	avgObjSizeString := fmt.Sprintf("%v", stats["avgObjSize"])
	indexesNumberString := fmt.Sprintf("%v", stats["nindexes"])

	capped := false

	if cappedString == "false" {
		capped = true
	}

	var count int64
	var size int64
	var avgObjSize int64
	var indexesNumber int64

	if i, err := strconv.ParseInt(countString, 10, 64); err == nil {
		count = i
	}

	if i, err := strconv.ParseInt(avgObjSizeString, 10, 64); err == nil {
		avgObjSize = i
	}

	if i, err := strconv.ParseInt(indexesNumberString, 10, 64); err == nil {
		indexesNumber = i
	}

	if i, err := strconv.ParseInt(sizeString, 10, 64); err == nil {
		size = i
	}

	return model.NewCollectionStats(count, avgObjSize, indexesNumber, capped, ns, size), nil
}
