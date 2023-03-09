package models

type User struct {
	Model
	FirstName string `json:"firstName" bson:"first_name" validate:"required"`
	LastName  string `json:"lastName" bson:"last_name" validate:"required"`
	EMail     string `json:"email" bson:"email" validate:"required,email"`
	Address
}
