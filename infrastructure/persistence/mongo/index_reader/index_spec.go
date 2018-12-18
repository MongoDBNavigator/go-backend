package index_reader

import "github.com/mongodb/mongo-go-driver/bson"

// Structure to represent index species
type indexSpec struct {
	Name       string        `bson:"name,omitempty"`
	Unique     bool          `bson:"unique,omitempty"`
	Background bool          `bson:"background,omitempty"`
	Sparse     bool          `bson:"sparse,omitempty"`
	Key        bson.Document `bson:"key,omitempty"`
}
