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
	// create instance, tmp dir
	t = new(TestDB)
	t.dir, err = ioutil.TempDir("", "mongo_testdb")
	if err != nil {
		t = nil
		return nil, err
	}
	// move to tmp
	err = os.Chdir(t.dir)
	if err != nil {
		return nil, err
	}
	// start db
	err = t.startDB()
	if err != nil {
		return nil, err
	}
	return
}

func (t *TestDB) startDB() (err error) {
	cmd := ("mongod --nohttpinterface --noprealloc --nojournal ")
	cmd += "--smallfiles --nssize=1 --oplogSize=1 --dbpath "
	cmd += t.dir + " --bind_ip=127.0.0.1 --port 40001"
	err = run(cmd)
	if err != nil {
		return err
	}
	return
}

// Close the test database, removing all temporary files.
func (t *TestDB) Close() error {
	err := t.rmTestDir()
	return err
}

func (t *TestDB) rmTestDir() (err error) {
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
	return
}

func run(command string) (err error) {
	// TODO add supervisord command execution
	return
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
