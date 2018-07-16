package model

// Model for database
type Database struct {
	name              string
	storageSize       int
	indexesNumber     int
	collectionsNumber int
}

// Getter for collectionsNumber
func (d *Database) CollectionsNumber() int {
	return d.collectionsNumber
}

// Getter for indexesNumber
func (d *Database) IndexesNumber() int {
	return d.indexesNumber
}

// Getter for storageSize
func (d *Database) StorageSize() int {
	return d.storageSize
}

// Getter for name
func (d *Database) Name() string {
	return d.name
}

// Constructor for database model
func NewDatabase(name string, storageSize int, indexesNumber int, collectionsNumber int) *Database {
	return &Database{
		name:              name,
		storageSize:       storageSize,
		indexesNumber:     indexesNumber,
		collectionsNumber: collectionsNumber,
	}
}
