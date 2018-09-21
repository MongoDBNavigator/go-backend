package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Fetch documents with pagination and filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadAll(conditions *value.ReadAllDocConditions) ([]interface{}, error) {
	return nil, nil
}
