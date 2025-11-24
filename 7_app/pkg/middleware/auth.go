package middleware

import (
	"api/configs"
	"api/pkg/jwt"
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		fmt.Println(isValid)
		fmt.Println(data)

		next.ServeHTTP(w, r)
	})
}
