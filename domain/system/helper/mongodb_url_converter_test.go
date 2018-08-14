package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoDBUrlConverterIp(t *testing.T) {
	url := "127.0.0.1:27017"

	MongoDBUrlConverter(url)

	assert.Equal(t, url, MongoDBUrlConverter(url))
}

func TestMongoDBUrlConverterCluster(t *testing.T) {
	url := "mongodb://admin:admin@gcp.mongodb.net:27017,gcp.mongodb.net:27017,gcp.mongodb.net:27017/test?&replicaSet=Test0&authSource=admin"

	MongoDBUrlConverter(url)

	assert.Equal(t, "gcp.mongodb.net:27017,gcp.mongodb.net:27017,gcp.mongodb.net:27017", MongoDBUrlConverter(url))
}
