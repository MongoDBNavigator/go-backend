package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Get documents count by filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadCount(conditions *value.ReadAllDocConditions) (int, error) {
	return 0, nil
}
