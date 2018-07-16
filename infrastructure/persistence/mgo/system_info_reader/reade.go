package system_info_reader

import "github.com/MongoDBNavigator/go-backend/domain/system/model"

// Get information about server (processor architecture, mongodb version, etc.)
func (rcv *systemInfoReader) Reade() (*model.SystemInfo, error) {
	buildInfo, err := rcv.db.BuildInfo()

	if err != nil {
		return nil, err
	}

	return model.NewSystemInfo(buildInfo.Version, buildInfo.Bits, rcv.url), nil
}
