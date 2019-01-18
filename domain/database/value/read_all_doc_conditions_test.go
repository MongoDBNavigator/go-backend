package value

import (
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/stretchr/testify/assert"
)

func TestNewReadAllDocConditions(t *testing.T) {
	dbName := DBName("DB")
	collName := CollName("COLLECTION")
	limit := 10
	skip := 2
	sort := []string{"-name"}
	filter := bson.M{"name": "test"}

	conditions := NewReadAllDocConditions(dbName, collName, limit, skip, sort, filter)

	assert.Equal(t, dbName, conditions.DbName())
	assert.Equal(t, collName, conditions.CollName())
	assert.Equal(t, limit, conditions.Limit())
	assert.Equal(t, skip, conditions.Skip())
	assert.Equal(t, sort, conditions.Sort())
	assert.Equal(t, filter, conditions.Filter())
}
