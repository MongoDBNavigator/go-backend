package representation

type Documents struct {
	Objects []interface{} `json:"objects"`
	Total   int64         `json:"total"`
}
