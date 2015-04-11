package sprinter_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"gopkg.in/mgo.v2"
)

// TestDB is a structure for the creation of a test mongo database.
type TestDB struct {
	session *mgo.Session
	dir     string
}

var (
	errInvalidTestDB = errors.New("invalid test database")
)

// Create  a new test database.
func NewTestDB() (t *TestDB, err error) {
	t = new(TestDB)
	t.dir, err = ioutil.TempDir("", "mongo_testdb")
	if err != nil {
		t = nil
		return
	}
	return
}

// Close the test database, removing all temporary files.
func (t *TestDB) Close() error {
	var err error

	if len(t.dir) > 0 {
		s := strings.Split(path.Dir(t.dir), string(os.PathSeparator))
		if len(s) < 2 {
			return errors.New("invalid test database directory")
		}
		base := "/" + s[1]
		if base != os.TempDir() {
			msg := "could not remove temp directory"
			msg += fmt.Sprintf("%s in %s", t.dir, base)
			return errors.New(msg)
		}
		err = os.RemoveAll(t.dir)
	}
	return err
}

func TestNewTestDB(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	if db == nil {
		t.Fatal(errInvalidTestDB)
	}
	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTestDB_Close(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
}
