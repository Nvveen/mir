package storage

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type testDB struct {
	*MongoDB
}

type testDBError struct {
	err string
	cmd string
}

var (
	TestDB *testDB // singleton
)

func newTestDBError(err string, cmd string) testDBError {
	return testDBError{err, cmd}
}

func (e testDBError) Error() string {
	if len(e.cmd) > 0 {
		return "Test database: " + e.err + " - Command output:\n" + e.cmd
	} else {
		return "Test database: " + e.err
	}
}

// Using a MongoDB backend, export a testing database (done so we don't have
// to export the testDB struct).
func NewTestDBMongo(m *MongoDB) *testDB {
	return &testDB{m}
}

func (db *testDB) StartMongoTesting() (err error) {
	log.Println("starting testing mongo")
	// TODO if database is running, restart it
	// check for supervisor, if it doesn't exist, we can't do testing
	if !supervisorExists() {
		return newTestDBError("could not find the supervisor daemon", "")
	}
	// start supervisor/mongo with the script
	err = run("cd ../storage/mongo_test && ./run.sh start")
	if err != nil {
		return newTestDBError("could not start the test database", err.(testDBError).cmd)
	}
	// Open connection
	err = TestDB.OpenConnection()
	if err != nil {
		return newTestDBError("could not open the database connection: "+err.Error(), "")
	}
	return nil
}

func (db *testDB) Reset() (err error) {
	log.Println("reset testing mongo")
	s := db.CloneSession()
	dbs, err := s.DatabaseNames()
	if err != nil {
		return newTestDBError(err.Error(), "")
	}
	for i := range dbs {
		s.DB(dbs[i]).DropDatabase()
	}
	return nil
}

func (db *testDB) StopMongoTesting() (err error) {
	log.Println("stopping testing mongo")
	// stop supervisor/mongo
	err = run("cd ../storage/mongo_test && ./run.sh stop")
	if err != nil {
		return newTestDBError("failed to stop the testing database", err.(testDBError).cmd)
	}
	return nil
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
	var output []byte
	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", command).CombinedOutput()
	} else {
		output, err = exec.Command("/bin/sh", "-c", command).CombinedOutput()
	}
	if err != nil {
		return newTestDBError("failed to execute command "+command, string(output))
	}
	return nil
}
