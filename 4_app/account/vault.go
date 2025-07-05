package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type DB interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type VaultWithDB struct {
	Vault
	db DB
}

func NewVault(db DB) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err)
	}

	return &VaultWithDB{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDB) FindAccounts(str string, checker func(Account, string) bool) (accounts []Account) {
	for _, account := range vault.Accounts {
		if checker(account, str) {
			accounts = append(accounts, account)
		}
	}

	return
}

func (vault *VaultWithDB) DeleteAccountByUrl(url string) (isDeleted bool) {
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

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToByteSlice()
	if err != nil {
		fmt.Println(err)
	}

	vault.db.Write(data)
}
