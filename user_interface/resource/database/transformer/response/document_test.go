package response

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/stretchr/testify/assert"
)

func TestDocumentsToView(t *testing.T) {
	name := "MyDB"
	storageSize := 1
	indexesNumber := 2
	collectionsNumber := 3

	db := model.NewDatabase(name, storageSize, indexesNumber, collectionsNumber)

	dbs := make([]interface{}, 1)
	dbs[0] = db

	total := 100
	view := DocumentsToView(dbs, total)

	assert.Len(t, view.Objects, 1)
	assert.Equal(t, total, view.Total)
}
