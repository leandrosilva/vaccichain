package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// CalculateHash hashes over a given string content
func CalculateHash(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

// IsHashValid checks if a hash matches a given difficulty
func IsHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
