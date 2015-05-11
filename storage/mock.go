package storage

type MockStorage map[string]map[string]string

func NewMockStorage() *MockStorage {
	m := &MockStorage{
		"linkindex": map[string]string{},
		"urlindex":  map[string]string{},
		"wordindex": map[string]string{},
	}
	return m
}

func (m *MockStorage) CloseConnection() {
}

func (m *MockStorage) OpenConnection() (err error) {
	return nil
}

func (m *MockStorage) InsertRecord(key string, url string, collection string) (err error) {
	(*m)[collection][key] = url
	return nil
}

func (m *MockStorage) String() string {
	res := "{\n"
	for name, collection := range *m {
		res += name + ":\n"
		for key, val := range collection {
			res += "\t" + key + ": " + val + "\n"
		}
	}
	res += "}"
	return res
}
