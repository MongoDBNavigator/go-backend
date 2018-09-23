package index_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Returns an array of documents that describe the existing indexes on a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.getIndexes/#db.collection.getIndexes
func (rcv *indexReader) ReadAll(dbName value.DBName, collName value.CollName) ([]*model.Index, error) {
	//indexView := rcv.db.Database(string(dbName)).Collection(name.Value().StringValue()).Indexes()
	//indexesCursor, err := indexView.List(context.Background())
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//index := bson.NewDocument()
	//for indexesCursor.Next(context.Background()) {
	//	index.Reset()
	//
	//	if err := indexesCursor.Decode(index); err != nil {
	//		return nil, err
	//	}
	//
	//	fmt.Printf("%v\n", index)
	//}

	return nil, nil
}
