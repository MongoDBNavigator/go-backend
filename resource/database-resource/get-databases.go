package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
)

func (rcv *databaseResource) getDatabases(request *restful.Request, response *restful.Response) {
	databases, err := rcv.databasesRepository.GetListDatabases()

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteEntity(transformer.DatabasesToView(databases))
}
