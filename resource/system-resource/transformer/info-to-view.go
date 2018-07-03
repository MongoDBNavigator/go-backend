package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/system-resource/representation"
)

func InfoToView(info *model.SystemInfo) *representation.Info {
	return &representation.Info{
		Url:                   info.GetUrl(),
		Version:               info.GetVersion(),
		ProcessorArchitecture: info.GetProcessorArchitecture(),
	}
}
