package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCollection(t *testing.T) {
	name := "MyCollection"
	docNumber := 1
	indexesNumber := 2
	avgObjSize := 3
	size := 4

	collection := NewCollection(name, docNumber, indexesNumber, avgObjSize, size)

	assert.Equal(t, name, collection.Name())
	assert.Equal(t, docNumber, collection.DocNumber())
	assert.Equal(t, indexesNumber, collection.IndexesNumber())
	assert.Equal(t, avgObjSize, collection.AvgObjSize())
	assert.Equal(t, size, collection.Size())
}
