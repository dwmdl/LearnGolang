package account

import (
	"PurpleSchool/app-4/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewVault() *Vault {
	file, err := files.ReadFile(files.FileName)
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err)
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountByUrl(url string) (accounts []Account) {
	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}

	return
}

func (vault *Vault) DeleteAccountByUrl(url string) (isDeleted bool) {
	var accounts []Account

	for _, account := range vault.Accounts {
		if !strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
			continue
		}

		isDeleted = true
	}

	vault.Accounts = accounts
	vault.save()

	return
}

func (vault *Vault) ToByteSlice() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToByteSlice()
	if err != nil {
		fmt.Println(err)
	}

	files.WriteFile(data, files.FileName)
}
