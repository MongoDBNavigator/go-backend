package mongo_system_repository

import "github.com/MongoDBNavigator/go-backend/persistence/model"

//
// Get information about server (processor architecture, mongodb version, etc.)
//
func (rcv *systemRepository) GetInfo() (*model.SystemInfo, error) {
	buildInfo, err := rcv.db.BuildInfo()

	if err != nil {
		return nil, err
	}

	return model.NewSystemInfo(buildInfo.Version, buildInfo.Bits, rcv.url), nil
}
