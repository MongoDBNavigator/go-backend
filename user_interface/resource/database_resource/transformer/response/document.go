package response

import "github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"

func DocumentsToView(data []interface{}, total int) *representation.Documents {
	result := new(representation.Documents)
	result.Total = total
	result.Objects = make([]interface{}, len(data))

	for i, record := range data {
		result.Objects[i] = record
	}

	return result
}
