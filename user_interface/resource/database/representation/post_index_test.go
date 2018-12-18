package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostIndexUnMarshal(t *testing.T) {
	index := PostIndex{}

	err := json.Unmarshal([]byte(`{"name":"MyIndex","unique":true,"background":true,"sparse":true,"fields":["name"]}`), &index)

	assert.Nil(t, err)
	assert.EqualValues(t, "MyIndex", index.Name)
	assert.EqualValues(t, []string{"name"}, index.Fields)
	assert.True(t, index.Unique)
	assert.True(t, index.Background)
	assert.True(t, index.Sparse)
}
