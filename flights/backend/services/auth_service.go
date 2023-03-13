package services

import (
	"errors"

	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/jwt"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
)

func Login(credentials dto.LoginDTO) (string, error) {
	user, err := repos.FindUserByEmail(credentials.EMail)
	if err != nil {
		return "", err //errors.New("user with given email does not exist")
	}

	if !verifyPassword(user.Password, credentials.Password) {
		return "", errors.New("incorrect password")
	}
	token, err := jwt.GenerateToken(user.EMail, user.Role)
	if err != nil {
		return "", errors.New("could not generate user token")
	}

	return token, nil
}

func verifyPassword(dbPassword string, dtoPassword string) bool {
	return dbPassword == dtoPassword
}