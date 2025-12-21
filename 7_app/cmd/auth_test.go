package main

import (
	"api/internal/auth"
	"api/internal/user"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load("../.env.test")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a@a.ru",
		Password: "$2a$10$f59AArFa/dCuJBMQ36HsRO8.4xCyCPGX0NKyk8sPjlCovGw0IFzXe",
		Name:     "TestName",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a@a.ru").
		Delete(&user.User{})
}

func TestLogin(t *testing.T) {
	t.Run("Login success", func(t *testing.T) {
		// prepare part
		db := initDb()
		initData(db)

		// test
		ts := httptest.NewServer(App())
		defer ts.Close()

		data, _ := json.Marshal(&auth.LoginRequest{
			Email:    "a@a.ru",
			Password: "123",
		})

		res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
		if err != nil {
			t.Fatal()
		}
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected %d got %d", http.StatusOK, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("Cannot read response body")
		}
		var resData auth.LoginResponse
		err = json.Unmarshal(body, &resData)
		if err != nil {
			t.Fatal("Cannot unmarshal json body")
		}
		if resData.Token == "" {
			t.Fatal("Got empty token field")
		}

		t.Log(resData.Token)
		removeData(db)
	})

	t.Run("Login failed", func(t *testing.T) {
		// prepare part
		db := initDb()
		initData(db)

		// test
		ts := httptest.NewServer(App())
		defer ts.Close()

		data, _ := json.Marshal(&auth.LoginRequest{
			Email:    "a@a123.ru",
			Password: "1",
		})

		res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
		if err != nil {
			t.Fatal()
		}
		if res.StatusCode == http.StatusOK {
			t.Fatalf("Expected %d got %d", http.StatusBadRequest, res.StatusCode)
		}

		removeData(db)
	})
}
