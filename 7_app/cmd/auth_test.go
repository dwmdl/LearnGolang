package main

import (
	"api/internal/auth"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
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
}
