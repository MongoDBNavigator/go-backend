package repository

import "github.com/MongoDBNavigator/go-backend/persistence/model"

type SystemRepositoryInterface interface {
	ResetIndexCache()
	GetInfo() (*model.SystemInfo, error)
}
