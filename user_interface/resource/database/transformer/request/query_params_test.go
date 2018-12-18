package request

import (
	"net/http"
	"testing"

	"github.com/emicklei/go-restful"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestExtractSkipSuccess(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost/?skip=10", nil)
	req := restful.NewRequest(r)

	value, err := ExtractSkip(req)

	assert.Nil(t, err)
	assert.EqualValues(t, 10, value)
}

func TestExtractLimitSuccess(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost/?limit=10", nil)
	req := restful.NewRequest(r)

	value, err := ExtractLimit(req)

	assert.Nil(t, err)
	assert.EqualValues(t, 10, value)
}

func TestExtractSortSuccess(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost/?sort[]=-name&sort[]=gender", nil)
	req := restful.NewRequest(r)

	value, err := ExtractSort(req)

	assert.Nil(t, err)
	assert.EqualValues(t, []string{"-name", "+gender"}, value)
}

func TestExtractFilterSuccess(t *testing.T) {
	r, _ := http.NewRequest("GET", `http://localhost/?filter[]=name:John`, nil)
	req := restful.NewRequest(r)

	value, err := ExtractFilter(req)

	assert.Nil(t, err)
	assert.EqualValues(t, bson.M{"name": "John"}, value)
}
