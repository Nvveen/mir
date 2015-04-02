package sprinter

import (
	"time"

	"gopkg.in/mgo.v2"
)

// TODO comments overal

type ReverseIndex map[string][]string

type Database struct {
	Host, Port             string
	Database, TestDatabase string
	Username, Password     string
	session                *mgo.Session
}

const (
	MongoDBHost  = "127.0.0.1"
	MongoDBPort  = "27017"
	AuthDatabase = "test"
	AuthUsername = "guest"
	AuthPassword = "welcome"
	TestDatabase = "test"
)

func NewDatabase() (db *Database, err error) {
	db = &Database{
		Host:     MongoDBHost,
		Port:     MongoDBPort,
		Database: AuthDatabase,
		Username: AuthUsername,
		Password: AuthPassword,
	}
	return
}

func (db *Database) OpenConnection() (err error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{db.Host},
		Timeout:  60 * time.Second,
		Database: db.Database,
		Username: db.Username,
		Password: db.Password,
	}
	db.session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return
	}
	db.session.SetMode(mgo.Monotonic, true)
	return
}

func (db *Database) CloseConnection() (err error) {
	db.session.Close()
	return
}
