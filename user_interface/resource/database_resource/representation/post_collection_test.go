package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCollectionUnMarshal(t *testing.T) {
	coll := PostCollection{}

	err := json.Unmarshal([]byte(`{"name":"MyCollection"}`), &coll)

	assert.Nil(t, err)
	assert.EqualValues(t, "MyCollection", coll.Name)
}
