package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostValidationUnMarshal(t *testing.T) {
	validation := PostValidation{}

	err := json.Unmarshal([]byte(
		`{"validationLevel":"Error","validationAction":"Strict","properties":[{"name":"test","required":true}]}`,
	), &validation)

	assert.Nil(t, err)
	assert.EqualValues(t, "Error", validation.ValidationLevel)
	assert.EqualValues(t, "Strict", validation.ValidationAction)
	assert.Len(t, validation.Properties, 1)
	assert.EqualValues(t, "test", validation.Properties[0].Name)
	assert.EqualValues(t, true, validation.Properties[0].Required)
}
