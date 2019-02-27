package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMarshal(t *testing.T) {
	errorView := Error{
		Message: "test error",
	}

	data, err := json.Marshal(errorView)

	assert.Nil(t, err)
	assert.Equal(t, `{"message":"test error"}`, string(data))
}
