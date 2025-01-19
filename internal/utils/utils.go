package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("your-256-bit-secret") // Ganti dengan secret key yang kuat

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Hashing password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// check password match
func VerifyPassword(hashedPassword, password string) bool {
	return HashPassword(password) == hashedPassword
}

// to check what is password match with hash
func VerifyToken(tokenString string) (*Claims, error) {
	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// make sure signing that we use is HS256
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	//verify claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// to return token
func GenerateToken(userID string) (string, error) {
	// create claims
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //token
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// create token with signing HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signing token with secret key
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		log.Fatalf("tokenString %v", err)
		return "", err
	}
	return tokenString, nil
}
