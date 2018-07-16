package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCredentialsUnMarshal(t *testing.T) {
	credentials := PostCredentials{}

	err := json.Unmarshal([]byte(`{"username":"Roman","password":"123"}`), &credentials)

	assert.Nil(t, err)
	assert.Equal(t, "Roman", credentials.Username)
	assert.Equal(t, "123", credentials.Password)
}
