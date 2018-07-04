package model

type Index struct {
	name       string
	unique     bool
	background bool
	sparse     bool
	fields     []string
}

func (rcv *Index) Fields() []string {
	return rcv.fields
}

func (rcv *Index) Sparse() bool {
	return rcv.sparse
}

func (rcv *Index) Background() bool {
	return rcv.background
}

func (rcv *Index) Unique() bool {
	return rcv.unique
}

func (rcv *Index) Name() string {
	return rcv.name
}

func NewIndex(name string, unique bool, background bool, sparse bool, fields []string) *Index {
	return &Index{
		name:       name,
		unique:     unique,
		fields:     fields,
		background: background,
		sparse:     sparse,
	}
}
