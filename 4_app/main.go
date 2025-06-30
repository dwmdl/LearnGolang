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
	fmt.Printf("Welcome to %s \n\n", appName)

Menu:
	for {
		userMenuChoice := promptData([]string{
			"1. Create an account",
			"2. Find an account",
			"3. Delete an account",
			"4. Exit\n",
			"Select one of the items",
		})

		switch userMenuChoice {
		case "1":
			getNewAccountData(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getNewAccountData(vault *account.VaultWithDB) {
	userLogin := promptData([]string{"Enter your login: "})
	userURL := promptData([]string{"Enter URL: "})
	userPassword := promptData([]string{"Enter password: "})

	newAccount, err := account.NewAccount(userLogin, userPassword, userURL)
	if err != nil {
		panic(err)
	}

	vault.AddAccount(*newAccount)
	newAccount.OutputData()
}

func findAccount(vault *account.VaultWithDB) {
	urlAccount := promptData([]string{"\nEnter the URL for searching"})

	foundAccounts := vault.FindAccountByUrl(urlAccount)
	if len(foundAccounts) == 0 {
		output.PrintError("Account not found")
	}

	for _, acc := range foundAccounts {
		acc.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	urlAccount := promptData([]string{"Enter the URL for searching"})
	isDeleted := vault.DeleteAccountByUrl(urlAccount)

	if isDeleted {
		fmt.Println("Account was successfully deleted")
	} else {
		output.PrintError("Account not found")
	}
}

func promptData[T any](prompt []T) (userInput string) {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	fmt.Scanln(&userInput)
	return
}
