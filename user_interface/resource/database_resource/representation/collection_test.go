package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionMarshal(t *testing.T) {
	collection := Collection{
		Name:            "MyCollection",
		DocumentsNumber: 1,
		IndexesNumber:   2,
		AvgObjSize:      3,
	}

	data, err := json.Marshal(collection)

	assert.Nil(t, err)
	assert.Equal(t, `{"name":"MyCollection","documentsNumber":1,"indexesNumber":2,"avgObjSize":3}`, string(data))
}
