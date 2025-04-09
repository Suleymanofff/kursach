package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTKey = []byte("7FyL2pBkQ9vR6tXwA1zD8cG5jH4nM7qP3eS9vY2uV6xZ")

func GenerateJWT(userID int, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}
