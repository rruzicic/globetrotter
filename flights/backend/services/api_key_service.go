package services

import (
	"log"
	"time"

	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
)

func GenerateAPIKey(temporary bool) models.APIKey {
	return repos.GenerateAPIKey(temporary)
}

func APIKeyExpired(key models.APIKey) bool {
	return time.Now().Before(key.Expiration)
}

func AddAPIKeyToUser(user models.User, key models.APIKey) bool {
	// Overwrites the last key!
	user.ApiKey = key
	return repos.UpdateUser(user)
}

func FindUserByAPIKey(key models.APIKey) (*models.User, error) {
	return repos.FindUserByAPIKey(key)
}

func BuyTicketForFriend(flightId string, key string, numOfTicketsOptional ...int) error {
	friend, err := FindUserByAPIKey(models.APIKey{Key: key, Expiration: time.Now()})
	if err != nil {
		log.Print("Could not find friend with given key. Error: ", err)
		return err
	}

	if APIKeyExpired(friend.ApiKey) {
		log.Print("Friend's key has expired. Error: ", err)
		return err
	}

	return BuyTicket(flightId, friend.EMail, numOfTicketsOptional...)
}
