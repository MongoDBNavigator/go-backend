package representation

type Info struct {
	Url             string `json:"url"`
	Version         string `json:"version"`
	CpuArchitecture int    `json:"processorArchitecture"`
}
