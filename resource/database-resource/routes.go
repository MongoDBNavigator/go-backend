package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"github.com/MongoDBNavigator/go-backend/resource"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

type databaseResource struct {
	databasesRepository   repository.DatabasesRepositoryInterface
	documentsRepository   repository.DocumentsRepositoryInterface
	collectionsRepository repository.CollectionsRepositoryInterface
}

func (rcv *databaseResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	dbTags := []string{"Databases"}
	collectionsTags := []string{"Collections"}
	documentsTags := []string{"Documents"}

	ws.Path("/api/v1/databases").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").
		To(rcv.getDatabases).
		Doc("Get all databases.").
		Writes(representation.Databases{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Databases{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, dbTags))

	ws.Route(ws.POST("").
		To(rcv.createDatabase).
		Doc("Create new database.").
		Reads(representation.PostDatabase{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, dbTags))

	ws.Route(ws.DELETE("/{databaseName}").
		To(rcv.dropDatabase).
		Doc("Drop an existing database.").
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, dbTags))

	ws.Route(ws.GET("/{databaseName}/collections").
		To(rcv.getCollections).
		Doc("Get all collections in database.").
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Writes(representation.Collections{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Collections{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, collectionsTags))

	ws.Route(ws.POST("/{databaseName}/collections").
		To(rcv.createCollection).
		Doc("Create new collection.").
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Reads(representation.PostCollection{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, collectionsTags))

	ws.Route(ws.DELETE("/{databaseName}/collections/{collectionName}").
		To(rcv.dropCollection).
		Doc("Drop an existing collection.").
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
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Writes(representation.Documents{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.Documents{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.POST("/{databaseName}/collections/{collectionName}/documents").
		To(rcv.createDocument).
		Doc("Create new document.").
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.DELETE("/{databaseName}/collections/{collectionName}/documents/{documentId}").
		To(rcv.dropDocument).
		Doc("Drop an existing document.").
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
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.PathParameter("documentId", "Document ID").DataType("string")).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	ws.Route(ws.PUT("/{databaseName}/collections/{collectionName}/documents/{documentId}").
		To(rcv.updateDocument).
		Doc("Update an existing document.").
		Param(ws.PathParameter("databaseName", "Database name").DataType("string")).
		Param(ws.PathParameter("collectionName", "Collection name").DataType("string")).
		Param(ws.PathParameter("documentId", "Document ID").DataType("string")).
		Returns(http.StatusAccepted, http.StatusText(http.StatusAccepted), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.Error{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.Error{}).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), representation.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, documentsTags))

	container.Add(ws)
}

func NewDatabaseResource(
	databasesRepository repository.DatabasesRepositoryInterface,
	collectionsRepository repository.CollectionsRepositoryInterface,
	documentsRepository repository.DocumentsRepositoryInterface,
) resource.ResourceInterface {
	return &databaseResource{
		databasesRepository:   databasesRepository,
		collectionsRepository: collectionsRepository,
		documentsRepository:   documentsRepository,
	}
}
