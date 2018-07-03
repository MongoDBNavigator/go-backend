package mongo_databases_repository

import (
	"fmt"
	"strconv"

	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"gopkg.in/mgo.v2/bson"
)

//
// Returns a document that lists all databases and returns basic database statistics.
// https://docs.mongodb.com/manual/reference/command/listDatabases/
//
func (rcv *databasesRepository) GetListDatabases() ([]*model.Database, error) {
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

		var collections, indexesNumber, storageSize int64

		if i, err := strconv.ParseInt(indexesString, 10, 64); err == nil {
			collections = i
		}

		if i, err := strconv.ParseInt(storageSizeString, 10, 64); err == nil {
			storageSize = i
		}

		if i, err := strconv.ParseInt(collectionsString, 10, 64); err == nil {
			indexesNumber = i
		}

		result[i] = model.NewDatabase(name, collections, indexesNumber, storageSize)
	}

	return result, nil
}
