package value

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/stretchr/testify/assert"
)

func TestCollNameOk(t *testing.T) {
	callName := CollName("TEST")

	assert.Nil(t, callName.Valid())
}

func TestCollNameEmpty(t *testing.T) {
	callName := CollName("")

	valid := callName.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.EmptyCollName, valid)
}
