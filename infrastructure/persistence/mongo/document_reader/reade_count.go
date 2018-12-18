package document_reader

import (
	"context"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Get documents count by filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadCount(conditions *value.ReadAllDocConditions) (int, error) {
	count, err := rcv.db.
		Database(string(conditions.DbName())).
		Collection(string(conditions.CollName())).
		Count(context.Background(), nil)

	if err != nil {
		return 0, err
	}

	return int(count), nil
}
