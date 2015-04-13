package sprinter

type Storage interface {
	OpenConnection() error
	CloseConnection()
	InsertRecord(key string, url string, collection string) error
}
