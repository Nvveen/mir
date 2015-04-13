package sprinter_test

import (
	"errors"
	"fmt"
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

type TestDBError struct {
	err       string
	cmdOutput string
}

func NewTestDBError(err string, cmdOutput string) TestDBError {
	return TestDBError{err, cmdOutput}
}

func (e TestDBError) Error() string {
	if len(e.cmdOutput) > 0 {
		return fmt.Sprintf("Test Database: %s: %s\n", e.err, e.cmdOutput)
	} else {
		return fmt.Sprintf("Test Database: %s\n", e.err)
	}
}

var (
	errNoSupervisor            = errors.New("could not find the supervisor daemon")
	errInvalidTestingDirectory = errors.New("invalid testing directory")
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
		return NewTestDBError("could not start the test database",
			err.(TestDBError).cmdOutput)
	}
	return nil
}

func StopMongoTesting() (err error) {
	// stop supervisor/mongo
	err = run("cd mongo_test && ./run.sh stop")
	if err != nil {
		return NewTestDBError("failed to stop the testing database",
			err.(TestDBError).cmdOutput)
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

func run(command string) (err error) {
	var b []byte
	if runtime.GOOS == "windows" {
		b, err = exec.Command("cmd", "/C", command).CombinedOutput()
	} else {
		b, err = exec.Command("/bin/sh", "-c", command).CombinedOutput()
	}
	if err != nil {
		return NewTestDBError("failed to execute command", string(b))
	}
	return nil
}

func TestNewTestDB(t *testing.T) {
	_, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
}
