package dto

import "github.com/rruzicic/globetrotter/flights/backend/models"

type AddUserApiKeyDTO struct {
	UserMail string         `json:"userMail"`
	APIKey   models.API_Key `json:"APIKey"`
}
