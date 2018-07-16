package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabasesMarshal(t *testing.T) {
	db := &Database{
		Name:          "MyDB",
		Collections:   1,
		IndexesNumber: 2,
		StorageSize:   3,
	}

	dbs := new(Databases)
	dbs.Objects = append(dbs.Objects, db)

	data, err := json.Marshal(dbs)

	assert.Nil(t, err)
	assert.Equal(t, `{"objects":[{"name":"MyDB","collections":1,"indexesNumber":2,"storageSize":3}]}`, string(data))
}
