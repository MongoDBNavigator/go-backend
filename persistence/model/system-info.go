package model

type SystemInfo struct {
	version               string
	processorArchitecture int
	url                   string
}

func NewSystemInfo(version string, processorArchitecture int, url string) *SystemInfo {
	return &SystemInfo{
		url:                   url,
		version:               version,
		processorArchitecture: processorArchitecture,
	}
}

func (rcv *SystemInfo) GetVersion() string {
	return rcv.version
}

func (rcv *SystemInfo) GetProcessorArchitecture() int {
	return rcv.processorArchitecture
}

func (rcv *SystemInfo) GetUrl() string {
	return rcv.url
}
