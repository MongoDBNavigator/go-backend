package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for index name
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
type IndexName string

// Validation method
func (rcv IndexName) Valid() error {
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyIndexName
	}

	return nil
}
