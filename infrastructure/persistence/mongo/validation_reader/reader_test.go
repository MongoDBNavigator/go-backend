package validation_reader

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/stretchr/testify/assert"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*repository.ValidationReader)(nil), new(validationReader))
}
