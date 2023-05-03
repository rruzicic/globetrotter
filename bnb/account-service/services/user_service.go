package services

import (
	"github.com/rruzicic/globetrotter/bnb/account-service/models"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterHost(user models.User) models.User {
	user.Role = models.HostRole
	return repos.CreateUser(user)
}

func RegisterGuest(user models.User) models.User {
	user.Role = models.GuestRole
	return repos.CreateUser(user)
}

func GetAll() []models.User {
	return repos.GetAllUsers()
}

func GetByEmail(email string) (*models.User, error) {
	return repos.GetUserByEmail(email)
}

func GetById(id primitive.ObjectID) (*models.User, error) {
	return repos.GetUserById(id)
}

func UpdateUser(user models.User) bool {
	return repos.UpdateUser(user)
}

func DeleteUser(id primitive.ObjectID) bool {
	return false
}
