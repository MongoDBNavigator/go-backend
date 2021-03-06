package representation

// Structure to representante POST body for index
type PostIndex struct {
	Name                    string      `json:"name"`
	Unique                  bool        `json:"unique"`
	Background              bool        `json:"background"`
	Sparse                  bool        `json:"sparse"`
	Fields                  []string    `json:"fields"`
	PartialFilterExpression interface{} `json:"partialFilterExpression,omitempty"`
}
