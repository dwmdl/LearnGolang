package jwt

import "testing"

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"

	jwtService := NewJWT("b7d92fa7e0a32688a32620ca693d8c5b35713d3d2152995f4a057bba6855a31a")
	token, err := jwtService.Create(Data{
		Email: email,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s is not equals to %s", email, data.Email)
	}
}
