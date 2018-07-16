package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseMarshal(t *testing.T) {
	db := Database{
		Name:          "MyDB",
		Collections:   1,
		IndexesNumber: 2,
		StorageSize:   3,
	}

	data, err := json.Marshal(db)

	assert.Nil(t, err)
	assert.Equal(t, `{"name":"MyDB","collections":1,"indexesNumber":2,"storageSize":3}`, string(data))
}
