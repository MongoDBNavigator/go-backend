package transformer

import (
	"github.com/MongoDBNavigator/go-backend/domain/system/model"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system_resource/representation"
)

func InfoToView(info *model.SystemInfo) *representation.Info {
	return &representation.Info{
		Url:             info.Url(),
		Version:         info.Version(),
		CpuArchitecture: info.CpuArchitecture(),
	}
}
