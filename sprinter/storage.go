package sprinter

import (
	"time"

	"gopkg.in/mgo.v2"
)

// TODO change default consts to use ENV vars

type ReverseIndex map[string][]string

type Database struct {
	Host, Port             string
	Database, TestDatabase string
	Username, Password     string
	session                *mgo.Session
}

const (
	MongoDBHost  = "127.0.0.1" // Default host
	MongoDBPort  = "27017"     // Default port
	AuthDatabase = "test"      // Default database
	AuthUsername = "testuser"  // Default user
	AuthPassword = "welcome"   // Default password
	TestDatabase = "test"      // Default testing database
)

// Constructs a new Database object with the default values.
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

// Open a MongoDB connection.
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

// Close a MongoDB connection
func (db *Database) CloseConnection() (err error) {
	db.session.Close()
	return
}
