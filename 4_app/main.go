package main

import (
	"PurpleSchool/app-4/account"
	"PurpleSchool/app-4/files"
	"PurpleSchool/app-4/output"
	"fmt"
)

const appName = "Password Manager"

func main() {
	vault := account.NewVault(files.NewJsonDB(files.FileName))
	fmt.Printf("Welcome to %s \n", appName)

Menu:
	for {
		userMenuChoice := userMenu()

		switch userMenuChoice {
		case 1:
			getNewAccountData(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func userMenu() (userChoice int) {
	fmt.Println("Select one of the items")
	fmt.Println("1. Create an account")
	fmt.Println("2. Find an account")
	fmt.Println("3. Delete an account")
	fmt.Println("4. Exit")

	fmt.Scan(&userChoice)
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

func findAccount(vault *account.VaultWithDB) {
	urlAccount := promptData("Enter the URL for searching")

	foundAccounts := vault.FindAccountByUrl(urlAccount)
	if len(foundAccounts) == 0 {
		output.PrintError("Account not found")
	}

	for _, acc := range foundAccounts {
		acc.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	urlAccount := promptData("Enter the URL for searching")
	isDeleted := vault.DeleteAccountByUrl(urlAccount)

	if isDeleted {
		fmt.Println("Account was successfully deleted")
	} else {
		output.PrintError("Account not found")
	}
}

func promptData(prompt string) (userInput string) {
	fmt.Println(prompt)
	fmt.Scanln(&userInput)
	return
}
