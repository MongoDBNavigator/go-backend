package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexesMarshal(t *testing.T) {
	index := &Index{
		Name:       "MyIndex",
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	index.Fields = append(index.Fields, "name")

	indexes := new(Indexes)
	indexes.Objects = append(indexes.Objects, index)

	data, err := json.Marshal(indexes)

	assert.Nil(t, err)
	assert.Equal(t, `{"objects":[{"name":"MyIndex","unique":true,"background":true,"sparse":true,"fields":["name"]}]}`, string(data))
}
