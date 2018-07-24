package value

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/stretchr/testify/assert"
)

func TestIndexNameOk(t *testing.T) {
	indexName := IndexName("_ID_")

	assert.Nil(t, indexName.Valid())
}

func TestIndexNameEmpty(t *testing.T) {
	indexName := IndexName("")

	valid := indexName.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.EmptyIndexName, valid)
}
