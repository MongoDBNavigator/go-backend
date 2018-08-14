package model

// Model for index
type Index struct {
	name                    string
	unique                  bool
	background              bool
	sparse                  bool
	fields                  []string
	partialFilterExpression interface{}
}

// Getter for partialFilterExpression
func (rcv *Index) PartialFilterExpression() interface{} {
	return rcv.partialFilterExpression
}

// Getter for background
func (rcv *Index) Fields() []string {
	return rcv.fields
}

// Getter for background
func (rcv *Index) Sparse() bool {
	return rcv.sparse
}

// Getter for background
func (rcv *Index) Background() bool {
	return rcv.background
}

// Getter for unique
func (rcv *Index) Unique() bool {
	return rcv.unique
}

// Getter for name
func (rcv *Index) Name() string {
	return rcv.name
}

// Constructor for index model
func NewIndex(name string, unique bool, background bool, sparse bool, fields []string, partialFilterExpression interface{}) *Index {
	return &Index{
		name:                    name,
		unique:                  unique,
		fields:                  fields,
		background:              background,
		sparse:                  sparse,
		partialFilterExpression: partialFilterExpression,
	}
}
