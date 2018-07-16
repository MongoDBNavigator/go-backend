package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexMarshal(t *testing.T) {
	index := Index{
		Name:       "MyIndex",
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	index.Fields = append(index.Fields, "name")

	data, err := json.Marshal(index)

	assert.Nil(t, err)
	assert.Equal(t, `{"name":"MyIndex","unique":true,"background":true,"sparse":true,"fields":["name"]}`, string(data))
}
