package services

import (
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
)

func RegisterUser(user models.User) bool {
	if repos.CreateUser(user) {
		return true
	}
	return false
}

func GetAllUsers() []models.User {
	return repos.FindAllUsers()
}

func FindUserByMail(mail string) (*models.User, error) {
	user, err := repos.FindUserByMail(mail)

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

func AddUserAPIKey(user models.User, api_key models.API_Key) bool {
	return repos.AddUserAPIKey(user, api_key)
}

func DeleteUserAPIKey(api_key string) bool {
	return repos.DeleteUserAPIKey(api_key)
}
