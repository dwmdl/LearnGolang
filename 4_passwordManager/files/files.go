package files

import (
	"fmt"
	"os"
)

type JsonDB struct {
	fileName string
}

func NewJsonDB(file string) *JsonDB {
	return &JsonDB{
		fileName: file,
	}
}

const FileName = "data.vault"

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success write!")
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
