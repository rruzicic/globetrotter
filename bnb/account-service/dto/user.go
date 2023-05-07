package dto

import "github.com/rruzicic/globetrotter/bnb/account-service/models"

type RegisterUserDTO struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	EMail     string `json:"email" binding:"required, email"`
	Password  string `json:"password" binding:"required, min=5"`
	Country   string `json:"country"`
	Street    string `json:"street"`
	StreetNum string `json:"streetNum"`
	ZIPCode   int    `json:"zip"`
}

func RegisterUserDTOtoUser(userDto RegisterUserDTO) models.User {
	return models.User{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		EMail:     userDto.EMail,
		Password:  userDto.Password,
		Address: models.Address{
			Country:   userDto.Country,
			Street:    userDto.Street,
			StreetNum: userDto.StreetNum,
			ZIPCode:   userDto.ZIPCode,
		},
	}
}

type UpdateUserDTO struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	EMail     string `json:"email" binding:"required, email"`
	Password  string `json:"password" binding:"required, min=5"`
	Country   string `json:"country"`
	Street    string `json:"street"`
	StreetNum string `json:"streetNum"`
	ZIPCode   int    `json:"zip"`
}

func UpdateUserDTOtoUser(userDto UpdateUserDTO) models.User {
	return models.User{
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		EMail:     userDto.EMail,
		Password:  userDto.Password,
		Address: models.Address{
			Country:   userDto.Country,
			Street:    userDto.Street,
			StreetNum: userDto.StreetNum,
			ZIPCode:   userDto.ZIPCode,
		},
	}
}

type CredentialsDTO struct {
	EMail    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required, min=5"`
}
