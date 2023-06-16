package repos

import (
	"math/rand"
	"time"
)

func generateAPIKeyString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 30
	byte_key := make([]byte, length)
	for i := range byte_key {
		byte_key[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(byte_key)
}
