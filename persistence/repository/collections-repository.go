package repository

import "github.com/MongoDBNavigator/go-backend/persistence/model"

type CollectionsRepositoryInterface interface {
	DropCollection(databaseName string, collectionName string) error
	CreateIndex(databaseName string, collectionName string, index *model.Index) error
	DropIndex(databaseName string, collectionName string, indexName string) error
	GetIndexes(databaseName string, collectionName string) ([]*model.Index, error)
	GetNumberOfDocuments(databaseName string, collectionName string) (int64, error)
	GetStats(databaseName string, collectionName string) (*model.CollectionStats, error)
	Create(*CollectionInfo) error
	GetCollectionsByDatabase(databaseName string) ([]*model.Collection, error)
}

type CollectionInfo struct {
	DatabaseName   string
	Name           string
	DisableIdIndex bool
	ForceIdIndex   bool
	Capped         bool
	MaxBytes       int
	MaxDocs        int
}

type Index struct {
	Key    []string
	Name   string
	Unique bool
}
