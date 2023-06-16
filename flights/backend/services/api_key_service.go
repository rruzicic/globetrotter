package services

import (
	"time"

	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
)

func APIKeyExpired(key models.APIKey) bool {
	return time.Now().Before(key.Expiration)
}

func AddAPIKeyToUser(user models.User, key models.APIKey) bool {
	// Overwrites the last key!
	user.ApiKey = key
	return repos.UpdateUser(user)
}
