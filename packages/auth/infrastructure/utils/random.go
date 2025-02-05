package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandomString generates a random string of the given length.
func RandomString(length int) (string, error) {
	randomBytes := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		randomBytes[i] = charset[randomIndex.Int64()]
	}

	return string(randomBytes), nil
}
