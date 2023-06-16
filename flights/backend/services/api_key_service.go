package services

import (
	"time"

	"github.com/rruzicic/globetrotter/flights/backend/models"
)

func APIKeyExpired(key models.APIKey) bool {
	return time.Now().Before(key.Expiration)
}
