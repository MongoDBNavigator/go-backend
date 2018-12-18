package representation

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Structure to representante POST body for database
type PostDatabase struct {
	Name value.DBName `json:"name"`
}
