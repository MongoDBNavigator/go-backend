package value

type ReadAllDocConditions struct {
	dbName   DBName
	collName CollName
	limit    int
	skip     int
	sort     map[string]int
	filter   []byte
}

func (rcv *ReadAllDocConditions) DbName() DBName {
	return rcv.dbName
}

func (rcv *ReadAllDocConditions) CollName() CollName {
	return rcv.collName
}

func (rcv *ReadAllDocConditions) Sort() map[string]int {
	return rcv.sort
}

func (rcv *ReadAllDocConditions) Skip() int {
	return rcv.skip
}

func (rcv *ReadAllDocConditions) Limit() int {
	return rcv.limit
}

func (rcv *ReadAllDocConditions) Filter() []byte {
	return rcv.filter
}

func NewReadAllDocConditions(
	dbName DBName,
	collName CollName,
	limit int,
	skip int,
	sort map[string]int,
	filter []byte,
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
