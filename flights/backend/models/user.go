package models

const (
	UserRole  string = "USER"
	AdminRole string = "ADMIN"
)

type User struct {
	Model     `bson:",inline"`
	FirstName string `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string `json:"lastName" bson:"last_name" binding:"required"`
	EMail     string `json:"email" bson:"email" binding:"required,email"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
	Address
}
