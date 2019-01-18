package system_info_reader

import (
	"context"
	"errors"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/system/helper"
	"github.com/MongoDBNavigator/go-backend/domain/system/model"
)

// Get information about server (processor architecture, mongodb version, etc.)
func (rcv *systemInfoReader) Reade() (*model.SystemInfo, error) {
	databaseNames, err := rcv.db.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(databaseNames) == 0 {
		return nil, errors.New("no available databases")
	}

	buildInfoResult := rcv.db.Database(databaseNames[0]).RunCommand(
		context.Background(),
		bson.D{{"buildInfo", 1}},
	)

	if buildInfoResult.Err() != nil {
		log.Println(err)
		return nil, buildInfoResult.Err()
	}

	raw, err := buildInfoResult.DecodeBytes()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var version string
	var bits int

	if versionRaw, err := raw.LookupErr("version"); err == nil {
		version = versionRaw.StringValue()
	}

	if bitsRaw, err := raw.LookupErr("bits"); err == nil {
		bits = int(bitsRaw.Int32())
	}

	return model.NewSystemInfo(version, bits, helper.MongoDBUrlConverter(rcv.url)), nil
}
