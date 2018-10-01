package db

import (
	"time"

	"config"
	"logger"

	mgo "gopkg.in/mgo.v2"
)

var mdb *mgo.Database
var log = logger.GetLogger()
var session *mgo.Session

func Init() {
	var err error
	cfg := config.GetConfig()
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{cfg.GetString("mongo.addr")},
		Timeout:  60 * time.Second,
		Database: cfg.GetString("mongo.db"),
		Username: cfg.GetString("mongo.user"),
		Password: cfg.GetString("mongo.pass"),
	}
	session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	mdb = session.DB("")
}

func GetDB() *mgo.Database {
	return mdb
}

func CloseDB() {
	session.Close()
}
