package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostDatabaseUnMarshal(t *testing.T) {
	db := PostDatabase{}

	err := json.Unmarshal([]byte(`{"name":"MyDB"}`), &db)

	assert.Nil(t, err)
	assert.EqualValues(t, "MyDB", db.Name)
}
