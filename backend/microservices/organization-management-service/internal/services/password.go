package services

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

const (
	saltSize   = 32 // Size of the random salt in bytes
	hashKeyLen = 64 // Length of the final hashed password in bytes
)

func HashPassword(password, salt []byte) ([]byte, error) {
	hash, err := scrypt.Key(password, salt, 16384, 8, 1, hashKeyLen)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func ValidatePassword(existingPassword, enteredPassword, salt []byte) bool {
	hash, err := HashPassword(enteredPassword, salt)
	if err != nil {
		return false
	}
	return base64.StdEncoding.EncodeToString(hash) == base64.StdEncoding.EncodeToString(existingPassword)
}

func GenerateRandomSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
