package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Hashing password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// to check what is password match with hash
func VerifyPassword(hashedPassword, password string) bool {
	return HashPassword(password) == hashedPassword
}

// to return token
func GenerateToken() string {
	return time.Now().Format("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
}
