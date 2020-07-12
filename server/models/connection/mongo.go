package connection

import (
	"crypto/tls"
	"github.com/globalsign/mgo"
	"github.com/go-bongo/bongo"
	"log"
	"net"
	"os"
	"time"
)

var mongoConn *bongo.Connection

func Mongo() *bongo.Connection {
	if mongoConn == nil {
		mongoConn = newMongo()
	}
	return mongoConn
}

func newMongo() *bongo.Connection {
	config := &bongo.Config{
		ConnectionString: os.Getenv("MONGO_ADDRESS"),
		Database:         os.Getenv("MONGO_DATABASE"),
	}

	if os.Getenv("MONGO_SSL") == "true" {
		// REF: https://github.com/go-bongo/bongo/pull/11
		if dialInfo, err := mgo.ParseURL(config.ConnectionString); err != nil {
			log.Fatal(err)
		} else {
			config.DialInfo = dialInfo
		}

		tlsConfig := &tls.Config{}
		config.DialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
		config.DialInfo.Timeout = time.Second * 3
	}

	conn, err := bongo.Connect(config)
	if err != nil {
		panic(err)
	}

	return conn
}
