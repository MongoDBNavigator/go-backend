package model

// Model for system info
type SystemInfo struct {
	cpuArchitecture int
	version         string
	url             string
}

// Getter for url
func (s *SystemInfo) Url() string {
	return s.url
}

// Getter for version
func (s *SystemInfo) Version() string {
	return s.version
}

// Getter for cpu architecture
func (s *SystemInfo) CpuArchitecture() int {
	return s.cpuArchitecture
}

// Constructor for SystemInfo model
func NewSystemInfo(version string, cpuArchitecture int, url string) *SystemInfo {
	return &SystemInfo{
		url:             url,
		version:         version,
		cpuArchitecture: cpuArchitecture,
	}
}
