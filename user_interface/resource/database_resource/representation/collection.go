package representation

// Structure for collection json representation
type Collection struct {
	Name            string `json:"name"`
	DocumentsNumber int    `json:"documentsNumber"`
	IndexesNumber   int    `json:"indexesNumber"`
	AvgObjSize      int    `json:"avgObjSize"`
}
