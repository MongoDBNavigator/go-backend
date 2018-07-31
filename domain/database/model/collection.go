package model

// Model for collection
type Collection struct {
	name          string
	size          int
	docNumber     int
	avgObjSize    int
	indexesNumber int
}

// Getter for size
func (c *Collection) Size() int {
	return c.size
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
func NewCollection(name string, docNumber int, indexesNumber int, avgObjSize int, size int) *Collection {
	return &Collection{
		name:          name,
		size:          size,
		docNumber:     docNumber,
		avgObjSize:    avgObjSize,
		indexesNumber: indexesNumber,
	}
}
