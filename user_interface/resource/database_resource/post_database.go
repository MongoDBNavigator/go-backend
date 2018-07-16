package database_resource

import (
	"net/http"

	"fmt"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
	"github.com/emicklei/go-restful"
)

// Method to post database
func (rcv *databaseResource) postDatabase(req *restful.Request, res *restful.Response) {
	postRequest := new(representation.PostDatabase)

	if err := req.ReadEntity(&postRequest); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	fmt.Println(postRequest)

	if err := rcv.databaseWriter.Create(postRequest.Name); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
}
