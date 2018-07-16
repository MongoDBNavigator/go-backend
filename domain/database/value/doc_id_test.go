package value

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/stretchr/testify/assert"
)

func TestDocIdOk(t *testing.T) {
	docId := DocId("ID")

	assert.Nil(t, docId.Valid())
}

func TestDocIdEmpty(t *testing.T) {
	docId := DocId("")

	valid := docId.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.EmptyDocId, valid)
}
