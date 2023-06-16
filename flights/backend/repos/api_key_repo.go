package repos

import (
	"math/rand"
	"time"

	"github.com/rruzicic/globetrotter/flights/backend/models"
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

func GenerateAPIKey(temporary bool) models.APIKey {
	key := generateAPIKeyString()
	if temporary {
		return models.APIKey{Key: key, Expiration: time.Now().AddDate(0, 3, 0)}
	}

	// some magic to get the max time so that expiration is at the heat death of the universe
	return models.APIKey{Key: key, Expiration: time.Unix(-2208988800, 0).Add(1<<63 - 1)}
}
