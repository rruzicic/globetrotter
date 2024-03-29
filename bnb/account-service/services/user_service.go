package services

import (
	"errors"
	"log"

	"github.com/rruzicic/globetrotter/bnb/account-service/dto"
	grpcclient "github.com/rruzicic/globetrotter/bnb/account-service/grpc_client"
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
	token, err := jwt.GenerateToken(user.EMail, user.Role, user.Id.Hex())
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

func UpdateNotificationPreferences(id string, notificationList []string) error {
	hexId, err := primitive.ObjectIDFromHex(id)
	if(err != nil) {
		log.Println("User service: Could not convert id to hex!")
		return err
	}
	user, err := GetById(hexId)
	if(err != nil) {
		log.Println("User service: Could not find user with id: ", hexId)
		return err
	}
	user.WantedNotifications = notificationList;
	_, err = UpdateUser(*user)
	if err != nil {
		log.Println("User service: Could not update user!")
		return err
	}
	return nil
}

func DeleteUser(id primitive.ObjectID) error {
	user, err := GetById(id)
	if err != nil {
		return err
	}

	if user.Role == models.GuestRole {
		reservations, err := grpcclient.GetActiveReservationsByUser(id.Hex())
		if err != nil {
			return err
		}
		if len(reservations) == 0 {
			repos.DeleteUser(id)
			return nil
		}
		return errors.New("there are active reservations for user")
	}

	if user.Role == models.HostRole {
		reservations, err := grpcclient.GetFutureActiveReservationsByHost(id.Hex())
		if err != nil {
			return err
		}
		if len(reservations) == 0 {
			repos.DeleteUser(id)
			return nil
		}
		return errors.New("there are active reservations for user")
	}
	return errors.New("role name invalid")
}

func IncrementCancellationsCounter(userId primitive.ObjectID) (*models.User, error) {
	user, err := GetById(userId)
	if err != nil {
		return &models.User{}, err
	}
	user.CancellationsCounter++
	user, err = UpdateUser(*user)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func verifyPassword(dbPassword string, dtoPassword string) bool {
	return dbPassword == dtoPassword
}

func AddAPIKeyToUser(email string, key string) (*models.User, error) {
	return repos.AddAPIKeyToUser(email, key)
}
