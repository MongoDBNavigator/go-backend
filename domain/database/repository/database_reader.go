package repository

import "github.com/MongoDBNavigator/go-backend/domain/database/model"

// Interface for database reader
// https://martinfowler.com/eaaCatalog/repository.html
type DatabaseReader interface {
	// Fetch all databases from current connection
	ReadAll() ([]*model.Database, error)
}
