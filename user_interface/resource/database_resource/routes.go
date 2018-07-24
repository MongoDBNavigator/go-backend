package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
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
	recoverMiddleware middleware.Middleware
}

// Method to register resource
func (rcv *databaseResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Filter(rcv.jwtMiddleware.Handle)
	//ws.Filter(rcv.recoverMiddleware.Handle)

	dbTags := []string{"Databases"}
	collectionsTags := []string{"Collections"}
	documentsTags := []string{"Documents"}
	indexesTags := []string{"Indexes"}
	validationTags := []string{"validation"}

	ws.Path("/api/v1/databases").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").
		To(rcv.getDatabases).
		Doc("Get all databases.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Writes(representation.Databases{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Databases{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, dbTags))

	ws.Route(ws.POST("").
		To(rcv.postDatabase).
		Doc("Create new database.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Reads(representation.PostDatabase{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, dbTags))

	ws.Route(ws.DELETE("/{databaseName}").
		To(rcv.deleteDatabase).
		Doc("Drop an existing database.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, dbTags))

	ws.Route(ws.GET("/{databaseName}/collections").
		To(rcv.getCollections).
		Doc("Get all collections in database.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Writes(representation.Collections{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Collections{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, collectionsTags))

	ws.Route(ws.POST("/{databaseName}/collections").
		To(rcv.postCollection).
		Doc("Create new collection.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Reads(representation.PostCollection{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, collectionsTags))

	ws.Route(ws.DELETE("/{databaseName}/collections/{collectionName}").
		To(rcv.deleteCollection).
		Doc("Drop an existing collection.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, collectionsTags))

	ws.Route(ws.GET("/{databaseName}/collections/{collectionName}/documents").
		To(rcv.getDocuments).
		Doc("Get documents in collection (with pagination and filters).").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.QueryParameter("limit", "Limit").DataType("number").DefaultValue("10")).
		Param(ws.QueryParameter("skip", "Skip").DataType("number").DefaultValue("0")).
		Param(ws.QueryParameter("sort[]", "Sorts").DataType("string")).
		Param(ws.QueryParameter("filter[]", "Filters").DataType("string")).
		Writes(representation.Documents{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Documents{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.POST("/{databaseName}/collections/{collectionName}/documents").
		To(rcv.postDocument).
		Doc("Create new document.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.DELETE("/{databaseName}/collections/{collectionName}/documents/{documentId}").
		To(rcv.deleteDocument).
		Doc("Drop an existing document.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.PathParameter("documentId", "Document ID").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.GET("/{databaseName}/collections/{collectionName}/documents/{documentId}").
		To(rcv.getDocument).
		Doc("Get document by ID.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.PathParameter("documentId", "Document ID").DataType("string")).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.PUT("/{databaseName}/collections/{collectionName}/documents/{documentId}").
		To(rcv.putDocument).
		Doc("Update an existing document.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.PathParameter("documentId", "Document ID").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.GET("/{databaseName}/collections/{collectionName}/indexes").
		To(rcv.getIndexes).
		Doc("Get collection indexes.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Writes(representation.Documents{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Documents{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, indexesTags))

	ws.Route(ws.POST("/{databaseName}/collections/{collectionName}/indexes").
		To(rcv.postIndex).
		Doc("Create index.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Reads(representation.PostIndex{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, indexesTags))

	ws.Route(ws.DELETE("/{databaseName}/collections/{collectionName}/indexes/{indexName}").
		To(rcv.deleteIndex).
		Doc("Create index.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.PathParameter("indexName", "Index name").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, indexesTags))

	ws.Route(ws.GET("/{databaseName}/collections/{collectionName}/validation").
		To(rcv.getValidation).
		Doc("Get collection validation.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Writes(representation.Validation{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Validation{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, validationTags))

	ws.Route(ws.POST("/{databaseName}/collections/{collectionName}/validation").
		To(rcv.postValidation).
		Doc("Post collection validation.").
		Param(ws.HeaderParameter("Authorization", "Bearer authentication").DataType("string")).
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Reads(representation.PostIndex{}).
		Returns(http.StatusOK, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, validationTags))

	container.Add(ws)
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
	recoverMiddleware middleware.Middleware,
) resource.Resource {
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
		recoverMiddleware: recoverMiddleware,
	}
}
