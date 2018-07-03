package main

import (
	"log"

	"net/http"

	"fmt"

	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-collections-repository"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-databases-repository"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-documents-repository"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-system-repository"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource"
	"github.com/MongoDBNavigator/go-backend/resource/swagger-resource"
	"github.com/MongoDBNavigator/go-backend/resource/system-resource"
	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2"
)

const (
	defaultMongoUrl   = "127.0.0.1:27017"
	defaultApiAddress = ":8080"
)

func main() {
	mongoSession, err := mgo.Dial(defaultMongoUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer mongoSession.Close()

	databasesRepository := mongo_databases_repository.New(mongoSession)
	collectionsRepository := mongo_collections_repository.New(mongoSession)
	systemRepository := mongo_system_repository.New(mongoSession, defaultMongoUrl)
	recordsRepository := mongo_documents_repository.New(mongoSession)

	var wsContainer = restful.NewContainer()

	database_resource.NewDatabaseResource(databasesRepository, collectionsRepository, recordsRepository).Register(wsContainer)
	system_resource.NewSystemResource(systemRepository).Register(wsContainer)

	swagger_resource.NewSwaggerResource(fmt.Sprintf("http://localhost%s", defaultApiAddress)).Register(wsContainer)

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CookiesAllowed: false,
		Container:      wsContainer,
	}

	wsContainer.Filter(cors.Filter)
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	server := http.Server{Addr: defaultApiAddress, Handler: wsContainer}

	log.Fatal(server.ListenAndServe())
}
