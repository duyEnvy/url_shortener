package common

import (
	"crypto/rand"
	"math/big"
)

func GenerateShortCode() (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength := 6
	randomString := make([]byte, codeLength)

	for i := 0; i < codeLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString), nil
}
