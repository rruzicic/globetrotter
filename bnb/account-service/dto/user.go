package dto

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
