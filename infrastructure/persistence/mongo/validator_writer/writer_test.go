package validator_writer

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/stretchr/testify/assert"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*repository.ValidationWriter)(nil), new(validatorWriter))
}
