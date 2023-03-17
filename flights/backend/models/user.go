package models

type User struct {
	Model     `bson:",inline"`
	FirstName string  `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string  `json:"lastName" bson:"last_name" binding:"required"`
	EMail     string  `json:"email" bson:"email" binding:"required,email"`
	Password  string  `json:"password" bson:"password"`
	APIKey    API_Key `json:"apiKey" bson:"api_key"`
	Address
}
