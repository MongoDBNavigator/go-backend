package representation

// Documents structure to represent documents list in json
type Documents struct {
	Objects []interface{} `json:"objects"`
	Total   int           `json:"total"`
}
