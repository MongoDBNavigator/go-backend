package representation

import (
	"encoding/json"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtTokenMarshal(t *testing.T) {
	token := JwtToken{
		Token: "test",
	}

	data, err := json.Marshal(token)

	assert.Nil(t, err)
	assert.Equal(t, `{"token":"test"}`, string(data))
}
