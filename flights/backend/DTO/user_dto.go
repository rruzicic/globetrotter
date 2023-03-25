package dto

type LoginDTO struct {
	EMail    string `json:"email" binding:"required,email,min=4"`
	Password string `json:"password" binding:"required,min=5"`
}
