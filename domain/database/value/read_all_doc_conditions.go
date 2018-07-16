package value

import (
	"gopkg.in/mgo.v2/bson"
)

type ReadAllDocConditions struct {
	dbName   DBName
	collName CollName
	limit    int
	skip     int
	sort     []string
	filter   bson.M
}

func (rcv *ReadAllDocConditions) DbName() DBName {
	return rcv.dbName
}

func (rcv *ReadAllDocConditions) CollName() CollName {
	return rcv.collName
}

func (rcv *ReadAllDocConditions) Sort() []string {
	return rcv.sort
}

func (rcv *ReadAllDocConditions) Skip() int {
	return rcv.skip
}

func (rcv *ReadAllDocConditions) Limit() int {
	return rcv.limit
}

func (rcv *ReadAllDocConditions) Filter() bson.M {
	return rcv.filter
}

func NewReadAllDocConditions(
	dbName DBName,
	collName CollName,
	limit int,
	skip int,
	sort []string,
	filter bson.M,
) *ReadAllDocConditions {
	return &ReadAllDocConditions{
		dbName:   dbName,
		collName: collName,
		limit:    limit,
		skip:     skip,
		sort:     sort,
		filter:   filter,
	}
}
