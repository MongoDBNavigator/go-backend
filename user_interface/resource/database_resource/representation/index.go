package representation

// Structure for index json representation
type Index struct {
	Name       string   `json:"name"`
	Unique     bool     `json:"unique"`
	Background bool     `json:"background"`
	Sparse     bool     `json:"sparse"`
	Fields     []string `json:"fields"`
}
