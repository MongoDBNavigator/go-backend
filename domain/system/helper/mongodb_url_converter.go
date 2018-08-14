package helper

import (
	"strings"

	"gopkg.in/mgo.v2"
)

func MongoDBUrlConverter(url string) string {

	dialInfo, err := mgo.ParseURL(url)

	if err != nil {
		return ""
	}

	return strings.Join(dialInfo.Addrs, ",")
}
