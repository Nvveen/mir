package sprinter

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO change default consts to use ENV vars
// TODO enhance errors
// TODO extend testing to setup test database, while moving every config
// to env vars for actual usage when not testing.
// TODO have Go start mongod, and keep it running while testing, adding users
// and the like.

// A reverse index representation of a bson MongoDB object.
type ReverseIndex struct {
	ID  bson.ObjectId `bson:"_id,omitempty"`
	Key string        `bson:"key"`
	// URLs are a list for now, eventually each will be a (weighted) element
	// so a proper index ranking can be made
	URLs []string `bson:"urls"`
}

type Database struct {
	Host, Port         string
	Database           string
	Username, Password string
	session            *mgo.Session
}

const (
	MongoDBHost  = "127.0.0.1" // Default host
	MongoDBPort  = "27017"     // Default port
	AuthDatabase = "test"      // Default database
	AuthUsername = "testuser"  // Default user
	AuthPassword = "welcome"   // Default password
)

// Constructs a new Database object with the default values.
func NewDatabase() *Database {
	return &Database{
		Host:     MongoDBHost,
		Port:     MongoDBPort,
		Database: AuthDatabase,
		Username: AuthUsername,
		Password: AuthPassword,
	}
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
func (db *Database) CloseConnection() {
	db.session.Close()
}

// Insert a new index record into the database.
func (db *Database) InsertRecord(key string, url string, collection string) (err error) {
	sessionCopy := db.session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(db.Database).C(collection)

	change := bson.M{"$addToSet": bson.M{"urls": &url}}
	err = c.Update(bson.M{"key": key}, change)
	if err == mgo.ErrNotFound {
		err = c.Insert(ReverseIndex{Key: key, URLs: []string{url}})
		if err != nil {
			return
		}
		err = nil
	}
	return
}
