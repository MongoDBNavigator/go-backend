package representation

type PostIndex struct {
	Name       string   `json:"name"`
	Unique     bool     `json:"unique"`
	Background bool     `json:"background"`
	Sparse     bool     `json:"sparse"`
	Fields     []string `json:"fields"`
}
