package storage

import (
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

type MongoDBError string

func (e MongoDBError) Error() string {
	return "Mongo DB: " + string(e)
}

// Open a MongoDB connection.
func (m *MongoDB) OpenConnection() (err error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{m.Host + ":" + m.Port},
		Timeout:  60 * time.Second,
		Database: m.Database,
		Username: m.Username,
		Password: m.Password,
	}
	m.session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return err
	}
	m.session.SetMode(mgo.Monotonic, true)
	return err
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
			return err
		}
		err = nil
	}
	return err
}

func (m *MongoDB) CloneSession() (session *mgo.Session) {
	return m.session.Clone()
}
