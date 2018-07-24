package model

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/stretchr/testify/assert"
)

func TestValidationIndex(t *testing.T) {
	validationAction := value.ValidationAction("error")
	validationLevel := value.ValidationLevel("warning")
	properties := make([]*ValidationProperty, 1)

	properties[0] = NewValidationProperty("test", true, "string", []string{}, "", 0, 0, "", 0, 0)

	validation := NewValidation(validationLevel, validationAction, properties)

	assert.Equal(t, validationAction, validation.ValidationAction())
	assert.Equal(t, validationLevel, validation.ValidationLevel())
	assert.Len(t, validation.Properties(), 1)
}
