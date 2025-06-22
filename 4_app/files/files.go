package files

import (
	"fmt"
	"os"
)

const FileName = "password.json"

func WriteFile(content []byte, fileName string) {
	file, err := os.Create(fileName)
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

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
