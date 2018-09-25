package system_info_reader

import (
	"context"
	"errors"

	"github.com/MongoDBNavigator/go-backend/domain/system/helper"
	"github.com/MongoDBNavigator/go-backend/domain/system/model"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Get information about server (processor architecture, mongodb version, etc.)
func (rcv *systemInfoReader) Reade() (*model.SystemInfo, error) {
	databaseNames, err := rcv.db.ListDatabaseNames(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	if len(databaseNames) == 0 {
		return nil, errors.New("no available databases")
	}

	buildInfo, err := rcv.db.Database(databaseNames[0]).RunCommand(
		context.Background(),
		bson.NewDocument(bson.EC.Int32("buildInfo", 1)),
	)

	if err != nil {
		return nil, err
	}

	version, err := buildInfo.Lookup("version")

	if err != nil {
		return nil, err
	}

	bits, err := buildInfo.Lookup("bits")

	if err != nil {
		return nil, err
	}

	return model.NewSystemInfo(
		version.Value().StringValue(),
		int(bits.Value().Int32()),
		helper.MongoDBUrlConverter(rcv.url),
	), nil
}
