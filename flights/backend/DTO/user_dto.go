package dto

import "github.com/rruzicic/globetrotter/flights/backend/models"

type LoginDTO struct {
	EMail    string `json:"email" binding:"required,email,min=4"`
	Password string `json:"password" binding:"required,min=5"`
}

type RegisterUserDTO struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	EMail     string `json:"email" binding:"required,email,min=4"`
	Password  string `json:"password" binding:"required,min=5"`
	Country   string `json:"country"`
	Street    string `json:"street"`
	StreetNum string `json:"streetNum"`
	ZIPCode   int    `json:"zip"`
}

func RegisterUserDTOToUser(user RegisterUserDTO) models.User {
	return models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		EMail:     user.EMail,
		Password:  user.Password,
		Address: models.Address{
			Country:   user.Country,
			Street:    user.Street,
			StreetNum: user.StreetNum,
			ZIPCode:   user.ZIPCode,
		},
	}
}
