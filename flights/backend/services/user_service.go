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
