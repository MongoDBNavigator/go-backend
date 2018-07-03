package model

type CollectionStats struct {
	ns            string
	count         int64
	size          int64
	avgObjSize    int64
	capped        bool
	indexesNumber int64
}

func NewCollectionStats(count int64, avgObjSize int64, indexesNumber int64, capped bool, ns string, size int64) *CollectionStats {
	return &CollectionStats{
		count:         count,
		avgObjSize:    avgObjSize,
		indexesNumber: indexesNumber,
		capped:        capped,
		ns:            ns,
		size:          size,
	}
}

func (rcv *CollectionStats) GetCount() int64 {
	return rcv.count
}

func (rcv *CollectionStats) GetAvgObjSize() int64 {
	return rcv.avgObjSize
}

func (rcv *CollectionStats) GetIndexesNumber() int64 {
	return rcv.indexesNumber
}

func (rcv *CollectionStats) IsCapped() bool {
	return rcv.capped
}

func (rcv *CollectionStats) GetNS() string {
	return rcv.ns
}

func (rcv *CollectionStats) GetSize() int64 {
	return rcv.size
}
