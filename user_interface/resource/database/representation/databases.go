package representation

// Databases structure to represent databases list in json
type Databases struct {
	Objects []*Database `json:"objects"`
}
