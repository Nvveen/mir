package storage

// TODO fix errors

type Storage interface {
	OpenConnection() error
	CloseConnection()
	InsertRecord(key string, url string, collection string) error
}
