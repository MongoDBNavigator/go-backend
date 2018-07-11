package repository

import "gopkg.in/mgo.v2/bson"

type DocumentsRepositoryInterface interface {
	GetAll(conditions *GetListConditions) ([]interface{}, error)
	GetNumberOfDocuments(conditions *GetListConditions) (int, error)
	Drop(databaseName string, collectionName string, documentId string) error
	Create(databaseName string, collectionName string, document interface{}) error
	GetById(databaseName string, collectionName string, documentId string) (interface{}, error)
	Update(databaseName string, collectionName string, documentId string, document interface{}) error
}

type GetListConditions struct {
	databaseName   string
	collectionName string
	limit          int
	skip           int
	sort           []string
	filter         bson.M
}

func (rcv *GetListConditions) Sort() []string {
	return rcv.sort
}

func (rcv *GetListConditions) Skip() int {
	return rcv.skip
}

func (rcv *GetListConditions) Limit() int {
	return rcv.limit
}

func (rcv *GetListConditions) CollectionName() string {
	return rcv.collectionName
}

func (rcv *GetListConditions) DatabaseName() string {
	return rcv.databaseName
}

func (rcv *GetListConditions) Filter() bson.M {
	return rcv.filter
}

func NewGetListConditions(
	databaseName string,
	collectionName string,
	limit int,
	skip int,
	sort []string,
	filter bson.M,
) *GetListConditions {
	return &GetListConditions{
		databaseName:   databaseName,
		collectionName: collectionName,
		limit:          limit,
		skip:           skip,
		sort:           sort,
		filter:         filter,
	}
}
