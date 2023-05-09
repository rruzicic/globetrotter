package services

import (
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/repos"
)

func CreateAccommodation(accommodation models.Accommodation) error {
	if err := repos.CreateAccommodation(accommodation); err != nil {
		return err
	}

	return nil
}

func UpdateAccommodation(accommodation models.Accommodation) error {
	if err := repos.UpdateAccommodation(accommodation); err != nil {
		return err
	}
	return nil
}
