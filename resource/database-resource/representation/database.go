package representation

type Database struct {
	Name          string `json:"name"`
	Collections   int64  `json:"collections"`
	IndexesNumber int64  `json:"indexesNumber"`
	StorageSize   int64  `json:"storageSize"`
}
