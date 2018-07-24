package main

import (
	"fmt"
	"log"

	"net/http"

	"strconv"

	"github.com/MongoDBNavigator/go-backend/infrastructure/helper"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/collection_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/collection_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/database_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/database_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/document_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/document_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/index_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/index_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/system_info_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/validation_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/validator_writer"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth_resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/swagger_resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system_resource"
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
	//env := helper.GetVar("ENV", "dev")
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

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()
	defer mongoSession.Close()

	databaseReader := database_reader.New(mongoSession)
	databaseWriter := database_writer.New(mongoSession)

	collectionsReader := collection_reader.New(mongoSession)
	collectionsWriter := collection_writer.New(mongoSession)

	documentReader := document_reader.New(mongoSession)
	documentWriter := document_writer.New(mongoSession)

	indexReader := index_reader.New(mongoSession)
	indexWriter := index_writer.New(mongoSession)

	validationReader := validation_reader.New(mongoSession)
	validationWriter := validator_writer.New(mongoSession)

	systemReader := system_info_reader.New(mongoSession, defaultMongoUrl)

	var wsContainer = restful.NewContainer()

	jwtMiddleware := middleware.NewJwtMiddleware(password)
	recoverMiddleware := middleware.NewRecoverMiddleware()

	database_resource.NewDatabaseResource(
		databaseReader,
		databaseWriter,
		collectionsReader,
		collectionsWriter,
		documentReader,
		documentWriter,
		indexReader,
		indexWriter,
		validationReader,
		validationWriter,
		jwtMiddleware,
		recoverMiddleware,
	).Register(wsContainer)
	system_resource.NewSystemResource(systemReader, jwtMiddleware, recoverMiddleware).Register(wsContainer)
	auth_resource.NewAuthResource(username, password, jwtExp).Register(wsContainer)

	//if env != defaultEnv {
	swagger_resource.NewSwaggerResource(fmt.Sprintf("http://localhost%s", apiAddress)).Register(wsContainer)

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CookiesAllowed: false,
		Container:      wsContainer,
	}

	wsContainer.Filter(cors.Filter)
	wsContainer.Filter(wsContainer.OPTIONSFilter)
	//}

	server := http.Server{Addr: apiAddress, Handler: wsContainer}

	log.Fatal(server.ListenAndServe())
}
