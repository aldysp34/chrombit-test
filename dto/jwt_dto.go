package dto

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(claims JWTClaims) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims.Issuer = os.Getenv("APP_NAME")
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
