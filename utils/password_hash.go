package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password string) string {
	hash := sha256.New()
	return base64.StdEncoding.EncodeToString(hash.Sum([]byte(password)))
}
