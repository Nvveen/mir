package storage

// TODO enhance errors
// TODO add database/user creation

type Storage interface {
	OpenConnection() error
	CloseConnection()
	InsertRecord(key string, url string, collection string) error
}
