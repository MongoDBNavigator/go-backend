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

	collection := NewCollection(name, docNumber, indexesNumber, avgObjSize)

	assert.Equal(t, name, collection.Name())
	assert.Equal(t, docNumber, collection.DocNumber())
	assert.Equal(t, indexesNumber, collection.IndexesNumber())
	assert.Equal(t, avgObjSize, collection.AvgObjSize())
}
