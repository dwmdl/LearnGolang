package auth_test

import (
	"api/configs"
	"api/internal/auth"
	"api/internal/user"
	"api/pkg/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.Handler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}

	userRepo := user.NewRepository(&db.DB{
		DB: gormDb,
	})
	handler := auth.Handler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		Service: auth.NewService(userRepo),
	}
	return &handler, mock, nil
}

func TestHandlerLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a@a.ru", "$2a$10$f59AArFa/dCuJBMQ36HsRO8.4xCyCPGX0NKyk8sPjlCovGw0IFzXe")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "123",
	})
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("got %d expected %d", w.Code, http.StatusOK)
	}
}

func TestHandlerRegisterSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "a@a.ru",
		Password: "123",
		Name:     "Erick",
	})
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("got %d expected %d", w.Code, http.StatusCreated)
	}
}
