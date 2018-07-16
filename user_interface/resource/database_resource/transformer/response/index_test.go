package response

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/stretchr/testify/assert"
)

func TestIndexToView(t *testing.T) {
	name := "MyDB"
	unique := true
	background := true
	sparse := true
	fields := []string{"name"}

	view := IndexToView(model.NewIndex(name, unique, background, sparse, fields))

	assert.Equal(t, name, view.Name)
	assert.Equal(t, unique, view.Unique)
	assert.Equal(t, background, view.Background)
	assert.Equal(t, sparse, view.Sparse)
	assert.Equal(t, fields, view.Fields)
}

func TestIndexesToView(t *testing.T) {
	name := "MyDB"
	unique := true
	background := true
	sparse := true
	fields := []string{"name"}

	index := model.NewIndex(name, unique, background, sparse, fields)

	indexes := make([]*model.Index, 1)
	indexes[0] = index

	view := IndexesToView(indexes)

	assert.Len(t, view.Objects, 1)
}
