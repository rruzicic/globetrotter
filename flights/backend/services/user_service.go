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

func GetUserByEmail(email string) (models.User, error) {
	return repos.FindUserByEmail(email)
}
