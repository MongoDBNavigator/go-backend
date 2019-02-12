package main

import (
	"log"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"

	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo"

	"net/http"

	"strconv"

	"github.com/MongoDBNavigator/go-backend/infrastructure/helper"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/collection_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/collection_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/database_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/database_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/document_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/document_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/index_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/index_writer"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/system_info_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/validation_reader"
	"github.com/MongoDBNavigator/go-backend/infrastructure/persistence/mongo/validator_writer"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/auth"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system"
)

const (
	defaultJwtExp          = "24" // hours
	defaultMongoUrl        = "mongodb://127.0.0.1:27017"
	defaultUsername        = "admin"
	defaultPassword        = "admin"
	defaultApiAddress      = ":8080"
	defaultStaticFilesPath = "/var/www"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	apiAddress := helper.GetVar("MN_PORT", defaultApiAddress)
	username := helper.GetVar("MN_USERNAME", defaultUsername)
	password := helper.GetVar("MN_PASSWORD", defaultPassword)
	jwtExp, err := strconv.Atoi(helper.GetVar("MN_JWT_EXP", defaultJwtExp))

	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mongo.NewMongoDBClient(helper.GetVar("MN_MONGO_URL", defaultMongoUrl))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success connect to mongodb.")

	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from ", r)
		}
	}()

	databaseReader := database_reader.New(mongoClient)
	databaseWriter := database_writer.New(mongoClient)

	collectionsReader := collection_reader.New(mongoClient)
	collectionsWriter := collection_writer.New(mongoClient)

	documentReader := document_reader.New(mongoClient)
	documentWriter := document_writer.New(mongoClient)

	indexReader := index_reader.New(mongoClient)
	indexWriter := index_writer.New(mongoClient)

	validationReader := validation_reader.New(mongoClient)
	validationWriter := validator_writer.New(mongoClient)

	systemReader := system_info_reader.New(mongoClient, defaultMongoUrl)

	wsContainer := mux.NewRouter()
	wsContainer.Use(middleware.NewContentTypeMiddleware().Handle) // globally content type

	jwtMiddleware := middleware.NewJwtMiddleware(password)
	recoverMiddleware := middleware.NewRecoverMiddleware()

	system.NewSystemResource(systemReader, jwtMiddleware, recoverMiddleware).Register(wsContainer)
	auth.NewAuthResource(username, password, jwtExp).Register(wsContainer)
	database.NewDatabaseResource(
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

	// Route for js app
	wsContainer.Handle("/", http.FileServer(http.Dir(defaultStaticFilesPath)))

	log.Println("MongoDb Navigator server start listening.")
	log.Fatal(http.ListenAndServe(apiAddress, handlers.CORS()(wsContainer)))
}
