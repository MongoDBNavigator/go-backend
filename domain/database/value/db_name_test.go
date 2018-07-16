package value

import (
	"testing"

	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/stretchr/testify/assert"
)

func TestDBNameOk(t *testing.T) {
	dbName := DBName("myDB")

	assert.Nil(t, dbName.Valid())
}

func TestDBNameEmpty(t *testing.T) {
	dbName := DBName("")

	valid := dbName.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.EmptyDBName, valid)
}

func TestDBNameInvalid1(t *testing.T) {
	dbName := DBName("myDB$")

	valid := dbName.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.NotValidDBName, valid)
}

func TestDBNameInvalid2(t *testing.T) {
	dbName := DBName("my DB")

	valid := dbName.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.NotValidDBName, valid)
}

func TestDBNameInvalid3(t *testing.T) {
	dbName := DBName("my.DB")

	valid := dbName.Valid()

	assert.Error(t, valid)
	assert.Equal(t, err.NotValidDBName, valid)
}
