package representation

// Database structure to represent database json
type Database struct {
	Name          string `json:"name"`
	Collections   int    `json:"collections"`
	IndexesNumber int    `json:"indexesNumber"`
	StorageSize   int    `json:"storageSize"`
}
