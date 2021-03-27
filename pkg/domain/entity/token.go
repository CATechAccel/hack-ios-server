package entity

import "github.com/dgrijalva/jwt-go"

type Token struct {
	jwt *jwt.Token
}

func NewToken() *Token {
	t := jwt.New(jwt.SigningMethodHS256)
	return &Token{
		jwt: t,
	}
}

func (t *Token) SetClaim(key string, value interface{}) {
	claims := t.jwt.Claims.(jwt.MapClaims)
	claims[key] = value
}

func (t *Token) Sign() (string, error) {
	signedToken, err := t.jwt.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
