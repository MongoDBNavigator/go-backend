package value

import (
	"strings"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
)

// Custom string type for validation action
// https://docs.mongodb.com/manual/core/schema-validation/
type ValidationAction string

// Validation method
func (rcv ValidationAction) Valid() error {
	if len(strings.TrimSpace(string(rcv))) == 0 {
		return err.EmptyValidationAction
	}

	switch rcv {
	case
		"error",
		"warning":
		return nil
	}

	return err.InvalidValidationAction
}
