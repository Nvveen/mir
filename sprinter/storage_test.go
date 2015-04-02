package sprinter

import "testing"

// We need this for disabling potential long-lasting MongoDB requests
// because they fail to timeout
const run = true

func TestNewDatabase(t *testing.T) {
	_, err := NewDatabase()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabase_OpenConnection(t *testing.T) {
	if run {
		db, err := NewDatabase()
		if err != nil {
			t.Fatal(err)
		}
		err = db.OpenConnection()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDatabase_CloseConnection(t *testing.T) {
	if run {
		db, err := NewDatabase()
		if err != nil {
			t.Fatal(err)
		}
		err = db.OpenConnection()
		if err != nil {
			t.Fatal(err)
		}
		err = db.CloseConnection()
		if err != nil {
			t.Fatal(err)
		}
	}
}
