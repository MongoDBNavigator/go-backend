package repository

import "github.com/MongoDBNavigator/go-backend/domain/system/model"

// Interface for system info reader (version, etc.)
// https://martinfowler.com/eaaCatalog/repository.html
type SystemInfoReader interface {
	// Read cpu architecture, version, conn url, etc.
	Reade() (*model.SystemInfo, error)
}
