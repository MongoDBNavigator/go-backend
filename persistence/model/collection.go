package model

type Collection struct {
	name          string
	fullName      string
	docNumber     int64
	indexesNumber int64
	avgObjSize    int64
}

func NewCollection(name string, docNumber int64, indexesNumber int64, avgObjSize int64) *Collection {
	return &Collection{
		name:          name,
		docNumber:     docNumber,
		indexesNumber: indexesNumber,
		avgObjSize:    avgObjSize,
	}
}

func (rcv *Collection) GetName() string {
	return rcv.name
}

func (rcv *Collection) GetDocumentsNumber() int64 {
	return rcv.docNumber
}

func (rcv *Collection) GetIndexesNumber() int64 {
	return rcv.indexesNumber
}

func (rcv *Collection) GetAvgObjSize() int64 {
	return rcv.avgObjSize
}
