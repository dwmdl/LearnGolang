package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputData() {
	fmt.Printf("Your login: %s \n", acc.login)
	fmt.Printf("Your password: %s \n", acc.password)
	fmt.Printf("URL: %s \n", acc.url)
}

func (acc *account) generatePassword(passwordLength int) {
	password := make([]rune, passwordLength)
	passwordSymbols := makeRange(33, 126)

	for i := range password {
		password[i] = passwordSymbols[rand.IntN(len(passwordSymbols))]
	}

	acc.password = string(password)
}

func newAccount(login, password, urlString string) (*account, error) {
	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if len(password) == 0 {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}

func main() {
	userLogin := promptData("Enter your login: ")
	userURL := promptData("Enter URL: ")
	userPassword := promptData("Enter password: ")

	myAccount, err := newAccount(userLogin, userPassword, userURL)
	if err != nil {
		panic(err)
	}

	myAccount.outputData()
}

func promptData(prompt string) (userInput string) {
	fmt.Println(prompt)
	fmt.Scanln(&userInput)
	return
}

func makeRange(min, max int) []int32 {
	a := make([]int32, max-min+1)
	for index := range a {
		a[index] = int32(min + index)
	}

	return a
}
