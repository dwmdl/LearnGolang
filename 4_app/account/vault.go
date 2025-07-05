package account

import (
	"PurpleSchool/app-4/encrypter"
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
	db  DB
	enc encrypter.Encrypter
}

func NewVault(db DB, enc encrypter.Encrypter) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	var vault Vault
	data := enc.Decrypt(file)
	err = json.Unmarshal(data, &vault)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Found %d account(s)\n\n", len(vault.Accounts))

	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
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

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToByteSlice()
	if err != nil {
		fmt.Println(err)
	}

	encData := vault.enc.Encrypt(data)
	vault.db.Write(encData)
}

func (vault *Vault) ToByteSlice() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return data, nil
}
