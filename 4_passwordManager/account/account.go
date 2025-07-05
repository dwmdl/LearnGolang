package account

import (
	"PurpleSchool/passwordManager/utils"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login,omitempty"`
	Password  string    `json:"password,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	var passwordLength int

	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(password) == 0 {
		fmt.Println("Enter how many symbols there should be in your password: ")
		_, errScan := fmt.Scanln(&passwordLength)
		if errScan != nil || passwordLength <= 0 {
			return nil, errors.New("INVALID_PASSWORD_LENGTH")
		}
		newAcc.generatePassword(passwordLength)
	}

	return newAcc, nil
}

func (acc *Account) generatePassword(passwordLength int) {
	password := make([]rune, passwordLength)
	passwordSymbols := utils.MakeRange(33, 126)

	for i := range password {
		password[i] = passwordSymbols[rand.IntN(len(passwordSymbols))]
	}

	acc.Password = string(password)
}

func (acc *Account) OutputData() {
	fmt.Print("\nACCOUNT DATA")
	fmt.Printf("\nYour login: %s \n", acc.Login)
	fmt.Printf("Your password: %s \n", acc.Password)
	fmt.Printf("URL: %s \n\n", acc.Url)
}
