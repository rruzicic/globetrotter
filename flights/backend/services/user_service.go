package services

import (
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
)

func RegisterUser(user models.User) bool {
	user.Role = models.UserRole
	return repos.CreateUser(user)
}

func GetAllUsers() []models.User {
	return repos.FindAllUsers()
}

func FindUserByEmail(mail string) (*models.User, error) {
	user, err := repos.FindUserByEmail(mail)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByAPIKey(api_key string) (*models.User, error) {
	user, err := repos.FindUserByAPIKey(api_key)

	if err != nil {
		return nil, err
	}

	return user, nil
}
