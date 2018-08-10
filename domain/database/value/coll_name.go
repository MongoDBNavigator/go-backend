package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for collection name
// https://docs.mongodb.com/manual/reference/limits/#Restriction-on-Collection-Names
// https://docs.mongodb.com/manual/reference/limits/#Restrictions-on-Database-Names-for-Unix-and-Linux-Systems
type CollName string

// Validation method
func (rcv CollName) Valid() error {
	// cannot be an empty string (e.g. "").
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyCollName
	}
	// cannot contain the $.
	if strings.ContainsAny(string(rcv), `$`) {
		return err.NotValidCollName
	}
	// cannot begin with the system. prefix. (Reserved for internal use.)
	if strings.HasPrefix(string(rcv), "system.") {
		return err.SystemPrefixContainsCollName
	}

	return nil
}
