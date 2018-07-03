package system_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/resource/system-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/system-resource/transformer"
)

func (rcv *systemResource) getInfo(request *restful.Request, response *restful.Response) {
	info, err := rcv.systemRepository.GetInfo()

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteEntity(transformer.InfoToView(info))
}
