package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for validation level
// https://docs.mongodb.com/manual/core/schema-validation/
type ValidationLevel string

// Validation method
func (rcv ValidationLevel) Valid() error {
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyValidationLevel
	}

	switch rcv {
	case
		"off",
		"strict",
		"moderate":
		return nil
	}

	return err.InvalidValidationLevel
}
