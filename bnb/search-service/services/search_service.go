package services

import (
	"github.com/rruzicic/globetrotter/bnb/search-service/dto"
	"github.com/rruzicic/globetrotter/bnb/search-service/models"
)

func Search() {

}

func SearchAccommodation(searchAccommodationDTO dto.SearchAccommodationDTO) ([]models.Accommodation, error) {
	flights, err := repos.GetFlightBySearchParams(searchAccommodationDTO)

	if err != nil {
		return nil, err
	}

	return flights, nil
}