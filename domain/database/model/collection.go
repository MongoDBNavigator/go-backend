package model

// Model for collection
type Collection struct {
	name          string
	docNumber     int
	avgObjSize    int
	indexesNumber int
}

// Getter for indexesNumber
func (c *Collection) IndexesNumber() int {
	return c.indexesNumber
}

// Getter for avgObjSize
func (c *Collection) AvgObjSize() int {
	return c.avgObjSize
}

// Getter for docNumber
func (c *Collection) DocNumber() int {
	return c.docNumber
}

// Getter for name
func (c *Collection) Name() string {
	return c.name
}

// Constructor for collection model
func NewCollection(name string, docNumber int, indexesNumber int, avgObjSize int) *Collection {
	return &Collection{
		name:          name,
		docNumber:     docNumber,
		avgObjSize:    avgObjSize,
		indexesNumber: indexesNumber,
	}
}
