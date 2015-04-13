package storage

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO change default consts to use ENV vars

type MongoDB struct {
	Host, Port         string
	Database           string
	Username, Password string
	session            *mgo.Session
}

// A reverse index representation of a bson MongoDB object.
type ReverseIndex struct {
	ID  bson.ObjectId `bson:"_id,omitempty"`
	Key string        `bson:"key"`
	// URLs are a list for now, eventually each will be a (weighted) element
	// so a proper index ranking can be made
	URLs []string `bson:"urls"`
}

type MongoDBError struct {
	err string
	m   *MongoDB
}

func NewMongoDBError(err string, m MongoDB) error {
	errfmsg := "MongoDB: %s\n"
	errfmsg += "\tDatabase: %s\n"
	errfmsg += "\tHost: %s\n"
	errfmsg += "\tUsername: %s\n"
	errfmsg += "\tPassword: %s\n"
	return fmt.Errorf(errfmsg, err, m.Database, m.Host, m.Username, m.Password)
}

// Constructs a new Database object with the default values.
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

// Open a MongoDB connection.
func (m *MongoDB) OpenConnection() (err error) {
	if len(m.Host) == 0 || len(m.Database) == 0 {
		return NewMongoDBError("empty mongo db configuration", *m)
	}
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{m.Host},
		Timeout:  60 * time.Second,
		Database: m.Database,
		Username: m.Username,
		Password: m.Password,
	}
	m.session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return
	}
	m.session.SetMode(mgo.Monotonic, true)
	return
}

// Close a MongoDB connection
func (m *MongoDB) CloseConnection() {
	m.session.Close()
}

// Insert a new index record into the database.
func (m *MongoDB) InsertRecord(key string, url string, collection string) (err error) {
	sessionCopy := m.session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(m.Database).C(collection)

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
