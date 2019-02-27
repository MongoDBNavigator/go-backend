package database

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/MongoDBNavigator/go-backend/user_interface"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
)

type databaseResource struct {
	databaseReader    repository.DatabaseReader
	databaseWriter    repository.DatabaseWriter
	collectionsReader repository.CollectionReader
	collectionsWriter repository.CollectionWriter
	documentReader    repository.DocumentReader
	documentWriter    repository.DocumentWriter
	indexReader       repository.IndexReader
	indexWriter       repository.IndexWriter
	validationReader  repository.ValidationReader
	validationWriter  repository.ValidationWriter
	jwtMiddleware     middleware.Middleware
}

// Method to register resource
func (rcv *databaseResource) Register(r *mux.Router) {
	sr := r.PathPrefix("/api/v1/databases").Subrouter()
	sr.Use(rcv.jwtMiddleware.Handle)

	sr.HandleFunc("", rcv.getDatabases).
		Methods("GET").
		Name("get_databases")

	sr.HandleFunc("", rcv.postDatabase).
		Methods("POST").
		Name("post_database")

	sr.HandleFunc("/{databaseName}", rcv.deleteDatabase).
		Methods("DELETE").
		Name("delete_database")

	sr.HandleFunc("/{databaseName}/collections", rcv.getCollections).
		Methods("GET").
		Name("get_collections")

	sr.HandleFunc("/{databaseName}/collections", rcv.postCollection).
		Methods("POST").
		Name("post_collection")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}", rcv.deleteCollection).
		Methods("DELETE").
		Name("delete_collection")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/documents", rcv.getDocuments).
		Methods("GET").
		Name("get_documents")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/documents", rcv.postDocument).
		Methods("POST").
		Name("post_document")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/documents/{documentId}", rcv.putDocument).
		Methods("PUT").
		Name("put_document")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/documents/{documentId}", rcv.getDocument).
		Methods("GET").
		Name("get_document")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/documents/{documentId}", rcv.deleteDocument).
		Methods("DELETE").
		Name("delete_document")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/indexes", rcv.getIndexes).
		Methods("GET").
		Name("get_indexes")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/indexes", rcv.postIndex).
		Methods("POST").
		Name("post_indexes")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/indexes/{indexName}", rcv.deleteIndex).
		Methods("DELETE").
		Name("delete_indexes")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/validation", rcv.getValidation).
		Methods("GET").
		Name("get_validation")

	sr.HandleFunc("/{databaseName}/collections/{collectionName}/validation", rcv.putValidation).
		Methods("PUT").
		Name("put_validation")
}

// Constructor for swaggerResource
func NewDatabaseResource(
	databaseReader repository.DatabaseReader,
	databaseWriter repository.DatabaseWriter,
	collectionsReader repository.CollectionReader,
	collectionsWriter repository.CollectionWriter,
	documentReader repository.DocumentReader,
	documentWriter repository.DocumentWriter,
	indexReader repository.IndexReader,
	indexWriter repository.IndexWriter,
	validationReader repository.ValidationReader,
	validationWriter repository.ValidationWriter,
	jwtMiddleware middleware.Middleware,
) user_interface.WebService {
	return &databaseResource{
		databaseReader:    databaseReader,
		databaseWriter:    databaseWriter,
		collectionsReader: collectionsReader,
		collectionsWriter: collectionsWriter,
		documentReader:    documentReader,
		documentWriter:    documentWriter,
		indexReader:       indexReader,
		indexWriter:       indexWriter,
		validationReader:  validationReader,
		validationWriter:  validationWriter,
		jwtMiddleware:     jwtMiddleware,
	}
}

// write status & body to response
func (rcv *databaseResource) writeResponse(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err)
	}
}

// write status & body to response
func (rcv *databaseResource) writeErrorResponse(w http.ResponseWriter, status int, err error) {
	rcv.writeResponse(w, status, representation.Error{Message: err.Error()})
}
