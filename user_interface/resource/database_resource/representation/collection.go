package representation

// Collection structure to represent collection json
type Collection struct {
	Name            string `json:"name"`
	Size            int    `json:"size"`
	DocumentsNumber int    `json:"documentsNumber"`
	IndexesNumber   int    `json:"indexesNumber"`
	AvgObjSize      int    `json:"avgObjSize"`
}
