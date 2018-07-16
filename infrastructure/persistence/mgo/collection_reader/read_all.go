package collection_reader

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Retrieve information, i.e. the name and options, about the collections and views in a database.
// https://docs.mongodb.com/manual/reference/command/listCollections/
func (rcv *collectionReader) ReadAll(dbName value.DBName) ([]*model.Collection, error) {
	collectionNames, err := rcv.db.DB(string(dbName)).CollectionNames()

	if err != nil {
		return nil, err
	}

	result := make([]*model.Collection, len(collectionNames))

	for i, name := range collectionNames {
		var stats bson.M
		// The collStats command returns a variety of storage statistics for a given collection.
		// https://docs.mongodb.com/manual/reference/command/collStats/
		if err := rcv.db.DB(string(dbName)).Run(bson.D{{"collStats", &name}}, &stats); err != nil {
			return nil, err
		}

		countString := fmt.Sprintf("%v", stats["count"])
		avgObjSizeString := fmt.Sprintf("%v", stats["avgObjSize"])
		indexesNumberString := fmt.Sprintf("%v", stats["nindexes"])

		var count int
		var avgObjSize int
		var indexesNumber int

		if i, err := strconv.Atoi(countString); err == nil {
			count = i
		}

		if i, err := strconv.Atoi(avgObjSizeString); err == nil {
			avgObjSize = i
		}

		if i, err := strconv.Atoi(indexesNumberString); err == nil {
			indexesNumber = i
		}

		result[i] = model.NewCollection(name, count, avgObjSize, indexesNumber)
	}

	return result, nil
}
