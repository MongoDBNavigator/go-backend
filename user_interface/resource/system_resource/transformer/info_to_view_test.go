package transformer

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/system/model"
	"github.com/stretchr/testify/assert"
)

func TestInfoToView(t *testing.T) {
	version := "4.0.0"
	cpuArchitecture := 64
	url := "127.0.0.1:27017"

	systemInfo := model.NewSystemInfo(version, cpuArchitecture, url)

	view := InfoToView(systemInfo)

	assert.Equal(t, version, view.Version)
	assert.Equal(t, cpuArchitecture, view.CpuArchitecture)
	assert.Equal(t, url, view.Url)
}
