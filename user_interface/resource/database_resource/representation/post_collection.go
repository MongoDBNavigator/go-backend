package representation

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Structure to representante POST body for collection
type PostCollection struct {
	Name value.CollName `json:"name"`
}
