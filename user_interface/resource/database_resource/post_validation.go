package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
)

// Method to get post validation
func (rcv *databaseResource) postValidation(req *restful.Request, res *restful.Response) {

	res.WriteHeader(http.StatusCreated)
}
