package services

import (
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/repos"
)

func GetRecommendedAccommodations(user models.User) ([]models.Accommodation, error) {
	return repos.GetRecommendedAccommodations(user)
}

func InitDBData() {
	repos.InitDBData()
}

func DropDB() error {
	return repos.DropDB()
}

func LoadMockDBData() {
	repos.LoadMockData()
}
