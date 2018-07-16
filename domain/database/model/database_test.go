package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	name := "MyDB"
	storageSize := 1
	indexesNumber := 2
	collectionsNumber := 3

	db := NewDatabase(name, storageSize, indexesNumber, collectionsNumber)

	assert.Equal(t, name, db.Name())
	assert.Equal(t, storageSize, db.StorageSize())
	assert.Equal(t, indexesNumber, db.IndexesNumber())
	assert.Equal(t, collectionsNumber, db.CollectionsNumber())
}
