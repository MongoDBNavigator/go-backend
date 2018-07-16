package response

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/stretchr/testify/assert"
)

func TestDatabasesToView(t *testing.T) {
	name := "MyDB"
	storageSize := 1
	indexesNumber := 2
	collectionsNumber := 3

	view := DatabaseToView(model.NewDatabase(name, storageSize, indexesNumber, collectionsNumber))

	assert.Equal(t, name, view.Name)
	assert.Equal(t, storageSize, view.StorageSize)
	assert.Equal(t, indexesNumber, view.IndexesNumber)
	assert.Equal(t, collectionsNumber, view.Collections)
}

func TestDatabaseToView(t *testing.T) {
	name := "MyDB"
	storageSize := 1
	indexesNumber := 2
	collectionsNumber := 3

	db := model.NewDatabase(name, storageSize, indexesNumber, collectionsNumber)

	dbs := make([]*model.Database, 1)
	dbs[0] = db

	view := DatabasesToView(dbs)

	assert.Len(t, view.Objects, 1)
}
