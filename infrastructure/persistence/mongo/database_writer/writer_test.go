package database_writer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*repository.DatabaseWriter)(nil), new(databaseWriter))
}
