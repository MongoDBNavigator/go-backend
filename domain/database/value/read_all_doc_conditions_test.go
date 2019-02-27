package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReadAllDocConditions(t *testing.T) {
	dbName := DBName("DB")
	collName := CollName("COLLECTION")
	limit := 10
	skip := 2
	sort := map[string]int{"name": 1}
	filter := []byte("name")

	conditions := NewReadAllDocConditions(dbName, collName, limit, skip, sort, filter)

	assert.Equal(t, dbName, conditions.DbName())
	assert.Equal(t, collName, conditions.CollName())
	assert.Equal(t, limit, conditions.Limit())
	assert.Equal(t, skip, conditions.Skip())
	assert.Equal(t, sort, conditions.Sort())
	assert.Equal(t, filter, conditions.Filter())
}
