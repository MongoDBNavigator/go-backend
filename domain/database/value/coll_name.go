package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for collection name
// https://docs.mongodb.com/manual/reference/limits/#Restriction-on-Collection-Names
type CollName string

// Validation method
func (rcv CollName) Valid() error {
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyCollName
	}

	return nil
}
