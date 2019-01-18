package helper

import (
	"log"
	"strings"

	"github.com/mongodb/mongo-go-driver/x/network/connstring"
)

func MongoDBUrlConverter(url string) string {

	dialInfo, err := connstring.Parse(url)

	if err != nil {
		log.Println(err)
		return ""
	}

	return strings.Join(dialInfo.Hosts, ",")
}
