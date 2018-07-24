package value

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/stretchr/testify/assert"
)

func TestValidationActionOk(t *testing.T) {
	action := ValidationAction("error")

	assert.Nil(t, action.Valid())
}

func TestValidationActionEmpty(t *testing.T) {
	action := ValidationAction("")

	valid := action.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.EmptyValidationAction, valid)
}

func TestValidationActionInvalid(t *testing.T) {
	action := ValidationAction("test")

	valid := action.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.InvalidValidationAction, valid)
}
