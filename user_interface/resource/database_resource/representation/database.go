package representation

// Structure for database json representation
type Database struct {
	Name          string `json:"name"`
	Collections   int    `json:"collections"`
	IndexesNumber int    `json:"indexesNumber"`
	StorageSize   int    `json:"storageSize"`
}
