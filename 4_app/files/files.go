package files

import (
	"PurpleSchool/app-4/output"
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

const FileName = "password.json"

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		output.PrintError(err)
	}

	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		output.PrintError(err)
	}

	fmt.Println("Success write!")
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		output.PrintError(err)
		return nil, err
	}

	return data, nil
}
