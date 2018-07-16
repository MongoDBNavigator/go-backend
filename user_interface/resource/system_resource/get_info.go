package system_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system_resource/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system_resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *systemResource) getInfo(request *restful.Request, response *restful.Response) {
	info, err := rcv.systemInfoReader.Reade()

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteEntity(transformer.InfoToView(info))
}
