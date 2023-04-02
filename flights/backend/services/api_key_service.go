package services

import (
	"math/rand"
	"time"

	"github.com/rruzicic/globetrotter/flights/backend/models"
)

func CreateAPIKey() string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	key := make([]rune, 30)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}
	return string(key)
}

func CheckAPIKeyExpiration(api_key models.API_Key) bool {
	if (api_key.Duration == time.Time{}) {
		return false
	}

	if time.Now().After(api_key.Duration) {
		return true
	}

	return false
}
