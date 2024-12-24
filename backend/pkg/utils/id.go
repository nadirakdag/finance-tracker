package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID creates a random ID string of specified length
func GenerateID(length int) string {
	bytes := make([]byte, length/2)

	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}
