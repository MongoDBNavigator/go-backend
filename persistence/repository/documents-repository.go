package repository

type DocumentsRepositoryInterface interface {
	GetAll(conditions *GetListConditions) ([]interface{}, error)
	Create(databaseName string, collectionName string, record interface{}) error
	Update(databaseName string, collectionName string, recordId string, document interface{}) error
	Drop(databaseName string, collectionName string, documentId string) error
	GetById(databaseName string, collectionName string, documentId string) (interface{}, error)
}

type GetListConditions struct {
	databaseName   string
	collectionName string
	limit          int
	skip           int
	sort           []string
}

func NewGetListConditions(
	databaseName string,
	collectionName string,
	limit int,
	skip int,
	sort []string,
) *GetListConditions {
	return &GetListConditions{
		databaseName:   databaseName,
		collectionName: collectionName,
		limit:          limit,
		skip:           skip,
		sort:           sort,
	}
}

func (rcv *GetListConditions) GetDatabaseName() string {
	return rcv.databaseName
}

func (rcv *GetListConditions) GetCollectionName() string {
	return rcv.collectionName
}

func (rcv *GetListConditions) GetLimit() int {
	return rcv.limit
}

func (rcv *GetListConditions) GetSkip() int {
	return rcv.skip
}

func (rcv *GetListConditions) GetSort() []string {
	return rcv.sort
}
