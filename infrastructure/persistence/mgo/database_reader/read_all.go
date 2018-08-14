package database_reader

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
)

// Returns a document that lists all databases and returns basic database statistics.
// https://docs.mongodb.com/manual/reference/command/listDatabases/
func (rcv *databaseReader) ReadAll() ([]*model.Database, error) {
	databaseNames, err := rcv.db.DatabaseNames()

	if err != nil {
		return nil, err
	}

	result := make([]*model.Database, len(databaseNames))

	for i, name := range databaseNames {
		var stats bson.M

		if err := rcv.db.DB(name).Run(bson.D{{"dbStats", 1}}, &stats); err != nil {
			return nil, err
		}

		indexesString := fmt.Sprintf("%#v", stats["indexes"])
		storageSizeString := fmt.Sprintf("%#v", stats["storageSize"])
		collectionsString := fmt.Sprintf("%#v", stats["collections"])

		var collections, indexesNumber, storageSize int

		if i, err := strconv.Atoi(collectionsString); err == nil {
			collections = i
		}

		if i, err := strconv.Atoi(storageSizeString); err == nil {
			storageSize = i
		}

		if i, err := strconv.Atoi(indexesString); err == nil {
			indexesNumber = i
		}

		result[i] = model.NewDatabase(name, storageSize, indexesNumber, collections)
	}

	return result, nil
}
