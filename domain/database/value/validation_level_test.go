package value

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/stretchr/testify/assert"
)

func TestValidationLevelOk(t *testing.T) {
	level := ValidationLevel("off")

	assert.Nil(t, level.Valid())
}

func TestValidationLevelEmpty(t *testing.T) {
	level := ValidationLevel("")

	valid := level.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.EmptyValidationLevel, valid)
}

func TestValidationLevelInvalid(t *testing.T) {
	level := ValidationLevel("test")

	valid := level.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.InvalidValidationLevel, valid)
}
