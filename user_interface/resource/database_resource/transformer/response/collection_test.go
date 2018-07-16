package response

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/stretchr/testify/assert"
)

func TestCollectionToView(t *testing.T) {
	name := "MyCollection"
	docNumber := 1
	indexesNumber := 2
	avgObjSize := 3

	view := CollectionToView(model.NewCollection(name, docNumber, indexesNumber, avgObjSize))

	assert.Equal(t, name, view.Name)
	assert.Equal(t, docNumber, view.DocumentsNumber)
	assert.Equal(t, indexesNumber, view.IndexesNumber)
	assert.Equal(t, avgObjSize, view.AvgObjSize)
}

func TestCollectionsToView(t *testing.T) {
	name := "MyCollection"
	docNumber := 1
	indexesNumber := 2
	avgObjSize := 3

	call := model.NewCollection(name, docNumber, indexesNumber, avgObjSize)

	calls := make([]*model.Collection, 1)
	calls[0] = call

	view := CollectionsToView(calls)

	assert.Len(t, view.Objects, 1)
}
