package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email, username string) (tokenstring string, err error) {
	expTime := time.Now().Add(1 * time.Minute)

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SigningString(jwtKey)
	return
}
