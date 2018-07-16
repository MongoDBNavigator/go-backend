package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSystemInfo(t *testing.T) {
	version := "4.0.0"
	cpuArchitecture := 64
	url := "127.0.0.1:27017"

	systemInfo := NewSystemInfo(version, cpuArchitecture, url)

	assert.Equal(t, version, systemInfo.Version())
	assert.Equal(t, cpuArchitecture, systemInfo.CpuArchitecture())
	assert.Equal(t, url, systemInfo.Url())
}
