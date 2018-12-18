package representation

// Indexes structure to represent indexes list in json
type Indexes struct {
	Objects []*Index `json:"objects"`
}
