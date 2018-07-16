package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for document id
type DocId string

// Validation method
func (rcv DocId) Valid() error {
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyDocId
	}

	return nil
}
