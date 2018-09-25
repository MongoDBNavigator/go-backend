package main

import (
	"fmt"
	"log"

	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo"

	"net/http"

	"strconv"

	"github.com/MongoDBNavigator/go-backend/infrastructure/helper"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/document_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/document_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/index_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/mgo_session"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/validation_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mgo/validator_writer"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth_resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/swagger_resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system_resource"
	"github.com/emicklei/go-restful"

	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/collection_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/collection_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/database_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/database_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/index_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/system_info_reader"
)

const (
	defaultEnv             = "prod"
	defaultJwtExp          = "24" // hours
	defaultMongoUrl        = "mongodb://127.0.0.1:27017"
	defaultUsername        = "admin"
	defaultPassword        = "admin"
	defaultApiAddress      = ":8080"
	defaultStaticFilesPath = "/var/www"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	env := helper.GetVar("ENV", "dev")
	apiAddress := helper.GetVar("MN_PORT", defaultApiAddress)
	username := helper.GetVar("MN_USERNAME", defaultUsername)
	password := helper.GetVar("MN_PASSWORD", defaultPassword)
	jwtExp, err := strconv.Atoi(helper.GetVar("MN_JWT_EXP", defaultJwtExp))

	if err != nil {
		log.Fatal(err)
	}

	mongoSession, err := mgo_session.MongoDBSessionFactory(helper.GetVar("MN_MONGO_URL", defaultMongoUrl))

	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mongo.MongoDBClientFactory(helper.GetVar("MN_MONGO_URL", defaultMongoUrl))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success connect to mongodb.")

	defer mongoSession.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	databaseReader := database_reader.New(mongoClient)
	databaseWriter := database_writer.New(mongoClient)

	collectionsReader := collection_reader.New(mongoClient)
	collectionsWriter := collection_writer.New(mongoClient)

	documentReader := document_reader.New(mongoSession)
	documentWriter := document_writer.New(mongoSession)

	indexReader := index_reader.New(mongoClient)
	indexWriter := index_writer.New(mongoSession)

	validationReader := validation_reader.New(mongoSession)
	validationWriter := validator_writer.New(mongoSession)

	systemReader := system_info_reader.New(mongoClient, defaultMongoUrl)

	var wsContainer = restful.NewContainer()

	jwtMiddleware := middleware.NewJwtMiddleware(password)
	recoverMiddleware := middleware.NewRecoverMiddleware()

	system_resource.NewSystemResource(systemReader, jwtMiddleware, recoverMiddleware).Register(wsContainer)
	auth_resource.NewAuthResource(username, password, jwtExp).Register(wsContainer)
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

	if env != defaultEnv {
		swagger_resource.NewSwaggerResource(fmt.Sprintf("http://localhost%s", apiAddress)).Register(wsContainer)

		cors := restful.CrossOriginResourceSharing{
			AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			CookiesAllowed: false,
			Container:      wsContainer,
		}

		wsContainer.Filter(cors.Filter)
		wsContainer.Filter(wsContainer.OPTIONSFilter)
	}

	// Route for js app
	wsContainer.Handle("/", http.FileServer(http.Dir(defaultStaticFilesPath)))

	server := http.Server{Addr: apiAddress, Handler: wsContainer}

	log.Println("MongoDb Navigator server start listening.")

	log.Fatal(server.ListenAndServe())
}
