package services

import (
	"errors"

	"github.com/rruzicic/globetrotter/bnb/account-service/dto"
	"github.com/rruzicic/globetrotter/bnb/account-service/jwt"
	"github.com/rruzicic/globetrotter/bnb/account-service/models"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterHost(user models.User) (*models.User, error) {
	user.Role = models.HostRole
	return repos.CreateUser(user)
}

func RegisterGuest(user models.User) (*models.User, error) {
	user.Role = models.GuestRole
	return repos.CreateUser(user)
}

func LoginUser(credentials dto.CredentialsDTO) (string, error) {
	user, err := repos.GetUserByEmail(credentials.EMail)
	if err != nil {
		return "", err
	}
	if !verifyPassword(user.Password, credentials.Password) {
		return "", errors.New("incorrect password")
	}
	token, err := jwt.GenerateToken(user.EMail, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
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

func UpdateUser(user models.User) (*models.User, error) {
	return repos.UpdateUser(user)
}

func DeleteUser(id primitive.ObjectID) bool {
	// TODO: implement
	return false
}

func verifyPassword(dbPassword string, dtoPassword string) bool {
	return dbPassword == dtoPassword
}
