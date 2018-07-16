package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionsMarshal(t *testing.T) {
	collection := &Collection{
		Name:            "MyCollection",
		DocumentsNumber: 1,
		IndexesNumber:   2,
		AvgObjSize:      3,
	}

	collections := new(Collections)
	collections.Objects = append(collections.Objects, collection)

	data, err := json.Marshal(collections)

	assert.Nil(t, err)
	assert.Equal(t, `{"objects":[{"name":"MyCollection","documentsNumber":1,"indexesNumber":2,"avgObjSize":3}]}`, string(data))
}
