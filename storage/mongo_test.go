package storage_test

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"gopkg.in/mgo.v2"
)

// TODO don't need to have errors as separate variables

type TestDB struct {
	session *mgo.Session
}

var (
	errNoSupervisor            = errors.New("could not find the supervisor daemon")
	errInvalidTestingDirectory = errors.New("invalid testing directory")
	errStartTestDB             = errors.New("could not start the testing database")
	errStopTestDB              = errors.New("could not stop the testing database")
	errDatabaseConnection      = errors.New("could not connect to the testing database")
)

func StartMongoTesting() (err error) {
	// check for supervisor, if it doesn't exist, we can't do testing
	if !supervisorExists() {
		return errNoSupervisor
	}
	// start supervisor/mongo with the script
	err = run("cd mongo_test && ./run.sh start")
	if err != nil {
		return errStartTestDB
	}
	return nil
}

func StopMongoTesting() (err error) {
	// stop supervisor/mongo
	err = run("cd mongo_test && ./run.sh stop")
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	err := StartMongoTesting()
	if err != nil {
		panic(err)
	}

	// run tests
	ret := m.Run()
	err = StopMongoTesting()
	if err != nil {
		panic(err)
	}

	os.Exit(ret)
}

func NewTestDB() (t *TestDB, err error) {
	t = new(TestDB)

	// connect to mongo
	t.session, err = mgo.Dial("127.0.0.1:40001")
	if err != nil {
		return nil, errDatabaseConnection
	}

	return t, nil
}

func supervisorExists() bool {
	env := os.Getenv("PATH")
	if len(env) == 0 {
		return false
	}
	for _, p := range filepath.SplitList(env) {
		fn := p + "/supervisord"
		if _, err := os.Stat(fn); err == nil {
			return true
		}
	}

	return false
}

func run(command string) error {
	var err error
	if runtime.GOOS == "windows" {
		_, err = exec.Command("cmd", "/C", command).CombinedOutput()
	} else {
		_, err = exec.Command("/bin/sh", "-c", command).CombinedOutput()
	}
	if err != nil {
		return errStartTestDB
	}
	return nil
}

func TestNewTestDB(t *testing.T) {
	_, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
}
