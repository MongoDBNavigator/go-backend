package helper

import (
	"testing"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringIDToBJSONFromObjectID(t *testing.T) {
	id := "5c6ab099a5f49fb972286062"

	actual := ConvertStringIDToBJSON(id)
	expected, _ := primitive.ObjectIDFromHex(id)

	assert.Equal(t, expected, actual.Map()["_id"])
}

func TestConvertStringIDToBJSONFromInt(t *testing.T) {
	actual := ConvertStringIDToBJSON("123")

	assert.Equal(t, 123, actual.Map()["_id"])
}

func TestConvertStringIDToBJSONFromString(t *testing.T) {
	actual := ConvertStringIDToBJSON("index")

	assert.Equal(t, "index", actual.Map()["_id"])
}
