package sprinter

import (
	"errors"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// TODO add database/user creation
// TODO only copy over sessions, don't initiate new connections each time

// We need this for disabling potential long-lasting MongoDB requests
// because they fail to timeout
const run = true

var (
	errNrKeys = errors.New("Invalid number of keys")
	errNrURLs = errors.New("Invalid number of urls in key")
)

func makeDB(t *testing.T) *Database {
	db := NewDatabase()
	db.Database = "gotest"
	db.Username = "gotestuser"
	db.Password = "welcome"

	err := db.OpenConnection()
	if err != nil {
		t.Fatal(err)
	}

	err = db.InsertRecord("www", "http://www.google.com", "urlindex")
	if err != nil {
		t.Fatal(err)
	}
	err = db.InsertRecord("www", "http://www.liacs.nl", "urlindex")
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func (db *Database) cleanDB(t *testing.T) {
	collections, err := db.session.DB("gotest").CollectionNames()
	if err != nil {
		t.Fatal(err)
	}
	for i := range collections {
		db.session.DB("gotest").C(collections[i]).DropCollection()
	}
}

func TestDatabase_OpenConnection(t *testing.T) {
	if run {
		db := NewDatabase()
		db.Database = "gotest"
		db.Username = "gotestuser"
		db.Password = "welcome"
		err := db.OpenConnection()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDatabase_CloseConnection(t *testing.T) {
	if run {
		db := NewDatabase()
		db.Database = "gotest"
		db.Username = "gotestuser"
		db.Password = "welcome"
		err := db.OpenConnection()
		if err != nil {
			t.Fatal(err)
		}
		db.CloseConnection()
	}
}

func TestDatabase_InsertRecord(t *testing.T) {
	if run {
		db := makeDB(t)
		db.cleanDB(t)
		defer db.cleanDB(t)

		err := db.InsertRecord("www", "http://www.leidenuniv.nl", "urlindex")
		if err != nil {
			t.Fatal(err)
		}

		// Determine if the insertions have happened correctly
		sessionCopy := db.session.Copy()
		defer sessionCopy.Close()
		c := sessionCopy.DB(db.Database).C("urlindex")
		var results []ReverseIndex
		err = c.Find(bson.M{"key": "www"}).All(&results)
		if err != nil {
			t.Fatal(err)
		}
		if len(results) != 1 {
			t.Fatalf("%s: %#v\n", errNrKeys, results)
		}
		if len(results[0].URLs) != 1 {
			t.Fatalf("%s: %#v", errNrURLs, results)
		}
	}
}
