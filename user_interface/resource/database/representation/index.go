package representation

// Index structure to represent index json
type Index struct {
	Name                    string      `json:"name"`
	Unique                  bool        `json:"unique"`
	Background              bool        `json:"background"`
	Sparse                  bool        `json:"sparse"`
	Fields                  []string    `json:"fields"`
	PartialFilterExpression interface{} `json:"partialFilterExpression,omitempty"`
}
