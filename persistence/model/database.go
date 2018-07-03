package model

type Database struct {
	name              string
	collectionsNumber int64
	indexesNumber     int64
	storageSize       int64
}

func NewDatabase(name string, collectionsNumber int64, indexesNumber int64, storageSize int64) *Database {
	return &Database{
		name:              name,
		collectionsNumber: collectionsNumber,
		indexesNumber:     indexesNumber,
		storageSize:       storageSize,
	}
}

func (rcv *Database) GetName() string {
	return rcv.name
}

func (rcv *Database) GetCollectionsNumber() int64 {
	return rcv.collectionsNumber
}

func (rcv *Database) GetIndexesNumber() int64 {
	return rcv.indexesNumber
}

func (rcv *Database) GetStorageSize() int64 {
	return rcv.storageSize
}
