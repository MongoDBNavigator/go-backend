package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoDBUrlConverterIp(t *testing.T) {
	url := "mongodb://127.0.0.1:27017"

	assert.Equal(t, "127.0.0.1:27017", MongoDBUrlConverter(url))
}

func TestMongoDBUrlConverterCluster(t *testing.T) {
	url := "mongodb://admin:admin@gcp.mongodb.net:27017,gcp.mongodb.net:27017,gcp.mongodb.net:27017/test?&replicaSet=Test0&authSource=admin"

	assert.Equal(t, "gcp.mongodb.net:27017,gcp.mongodb.net:27017,gcp.mongodb.net:27017", MongoDBUrlConverter(url))
}
