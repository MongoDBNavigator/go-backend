package representation

// Structure for documents list json representation
type Documents struct {
	Objects []interface{} `json:"objects"`
	Total   int           `json:"total"`
}
