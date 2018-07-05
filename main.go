package main

import (
	"fmt"
	"log"

	"net/http"

	"strconv"

	"github.com/MongoDBNavigator/go-backend/helper"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-collections-repository"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-databases-repository"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-documents-repository"
	"github.com/MongoDBNavigator/go-backend/persistence/repository/mongo-system-repository"
	"github.com/MongoDBNavigator/go-backend/resource/auth-resource"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource"
	"github.com/MongoDBNavigator/go-backend/resource/middleware"
	"github.com/MongoDBNavigator/go-backend/resource/swagger-resource"
	"github.com/MongoDBNavigator/go-backend/resource/system-resource"
	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2"
)

const (
	defaultMongoUrl   = "127.0.0.1:27017"
	defaultApiAddress = ":8080"
	defaultUsername   = "admin"
	defaultPassword   = "admin"
	defaultJwtExp     = "24" // hours
	defaultEnv        = "prod"
)

func main() {
	env := helper.GetVar("ENV", defaultEnv)
	apiAddress := helper.GetVar("PORT", defaultApiAddress)
	username := helper.GetVar("USERNAME", defaultUsername)
	password := helper.GetVar("PASSWORD", defaultPassword)
	jwtExp, err := strconv.Atoi(helper.GetVar("JWT_EXP", defaultJwtExp))

	if err != nil {
		log.Fatal(err)
	}

	mongoSession, err := mgo.Dial(helper.GetVar("MONGO_URL", defaultMongoUrl))

	if err != nil {
		log.Fatal(err)
	}

	defer mongoSession.Close()

	databasesRepository := mongo_databases_repository.New(mongoSession)
	collectionsRepository := mongo_collections_repository.New(mongoSession)
	systemRepository := mongo_system_repository.New(mongoSession, defaultMongoUrl)
	recordsRepository := mongo_documents_repository.New(mongoSession)

	var wsContainer = restful.NewContainer()

	jwtMiddleware := middleware.NewJwtMiddleware(password)

	database_resource.NewDatabaseResource(databasesRepository, collectionsRepository, recordsRepository, jwtMiddleware).Register(wsContainer)
	system_resource.NewSystemResource(systemRepository, jwtMiddleware).Register(wsContainer)
	auth_resource.NewAuthResource(username, password, jwtExp).Register(wsContainer)

	if env != defaultEnv {
		swagger_resource.NewSwaggerResource(fmt.Sprintf("http://localhost%s", apiAddress)).Register(wsContainer)

		cors := restful.CrossOriginResourceSharing{
			AllowedHeaders: []string{"Content-Type", "Accept"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			CookiesAllowed: false,
			Container:      wsContainer,
		}

		wsContainer.Filter(cors.Filter)
		wsContainer.Filter(wsContainer.OPTIONSFilter)
	}

	server := http.Server{Addr: apiAddress, Handler: wsContainer}

	log.Fatal(server.ListenAndServe())
}
