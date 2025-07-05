package main

import (
	"PurpleSchool/app-4/account"
	"PurpleSchool/app-4/files"
	"fmt"
	"github.com/joho/godotenv"
	"strings"
)

const appName = "Password Manager"

var menu = map[string]func(*account.VaultWithDB){
	"1": getNewAccountData,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to found ENV file")
	}

	vault := account.NewVault(files.NewJsonDB(files.FileName))
	fmt.Printf("Welcome to %s \n\n", appName)

Menu:
	for {
		userMenuChoice := promptData(
			"1. Create an account",
			"2. Find an account by URL",
			"3. Find an account by Login",
			"4. Delete an account",
			"5. Exit\n",
			"Select one of the items",
		)

		menuFunc := menu[userMenuChoice]
		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)
	}
}

func promptData(prompt ...any) (userInput string) {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func getNewAccountData(vault *account.VaultWithDB) {
	userLogin := promptData("Enter your login: ")
	userURL := promptData("Enter URL: ")
	userPassword := promptData("Enter password: ")

	newAccount, err := account.NewAccount(userLogin, userPassword, userURL)
	if err != nil {
		panic(err)
	}

	vault.AddAccount(*newAccount)
	newAccount.OutputData()
}

func findAccountByURL(vault *account.VaultWithDB) {
	urlAccount := promptData("Enter the URL for searching")

	foundAccounts := vault.FindAccounts(urlAccount, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	outputFindingResult(&foundAccounts)
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData("Enter the login for searching")

	foundAccounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	outputFindingResult(&foundAccounts)
}

func outputFindingResult(foundAccounts *[]account.Account) {
	if len(*foundAccounts) == 0 {
		fmt.Println("Account not found")
	}

	for _, acc := range *foundAccounts {
		acc.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	urlAccount := promptData("Enter the URL for searching")
	isDeleted := vault.DeleteAccountByUrl(urlAccount)

	if isDeleted {
		fmt.Println("Account was successfully deleted")
	} else {
		fmt.Println("Account not found")
	}
}
