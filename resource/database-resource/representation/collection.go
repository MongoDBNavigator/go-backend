package representation

type Collection struct {
	Name            string `json:"name"`
	DocumentsNumber int64  `json:"documentsNumber"`
	IndexesNumber   int64  `json:"indexesNumber"`
	AvgObjSize      int64  `json:"avgObjSize"`
}
