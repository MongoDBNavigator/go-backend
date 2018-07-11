package representation

type Documents struct {
	Objects []interface{} `json:"objects"`
	Total   int           `json:"total"`
}
