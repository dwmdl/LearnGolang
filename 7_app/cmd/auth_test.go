package main

import (
	"api/internal/auth"
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
	err := godotenv.Load("cmd/.env.test")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestLogin(t *testing.T) {
	t.Run("Login success", func(t *testing.T) {
		// prepare part
		//db := initDb()

		// test
		ts := httptest.NewServer(App())
		defer ts.Close()

		data, _ := json.Marshal(&auth.LoginRequest{
			Email:    "a@a123.ru",
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
	})

	t.Run("Login failed", func(t *testing.T) {
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
	})
}
