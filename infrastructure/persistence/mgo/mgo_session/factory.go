package mgo_session

import (
	"log"

	"crypto/tls"
	"net"

	"gopkg.in/mgo.v2"
)

func MongoDBSessionFactory(url string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(url)

	if err != nil {
		log.Fatal(err)
	}

	if dialInfo.Password == "" || dialInfo.Username == "" {
		return mgo.Dial(url)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
		return conn, err
	}

	return mgo.DialWithInfo(dialInfo)
}
