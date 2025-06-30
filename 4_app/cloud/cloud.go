package cloud

type DB struct {
	url string
}

func NewCloudDB(urlToDB string) *DB {
	return &DB{
		url: urlToDB,
	}
}

func (db *DB) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *DB) Write(content []byte) {
	//
}
