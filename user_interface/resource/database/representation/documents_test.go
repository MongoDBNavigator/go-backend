package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentsMarshal(t *testing.T) {
	doc := &Database{
		Name:          "MyDB",
		Collections:   1,
		IndexesNumber: 2,
		StorageSize:   3,
	}

	docs := new(Documents)
	docs.Objects = append(docs.Objects, doc)
	docs.Total = len(docs.Objects)

	data, err := json.Marshal(docs)

	assert.Nil(t, err)
	assert.Equal(t, `{"objects":[{"name":"MyDB","collections":1,"indexesNumber":2,"storageSize":3}],"total":1}`, string(data))
}
