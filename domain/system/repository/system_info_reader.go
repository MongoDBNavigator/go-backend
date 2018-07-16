package repository

import "github.com/MongoDBNavigator/go-backend/domain/system/model"

type SystemInfoReader interface {
	Reade() (*model.SystemInfo, error)
}
