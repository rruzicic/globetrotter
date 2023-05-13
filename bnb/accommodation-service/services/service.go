package services

import (
	"log"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/dtos"
	grpcclient "github.com/rruzicic/globetrotter/bnb/accommodation-service/grpc_client"
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

func UpdatePriceInterval(updatePriceDTO dtos.UpdatePriceDTO) (bool, error) {
	accommodation, err := repos.GetAccommodationById(updatePriceDTO.AccommodationId)
	if err != nil {
		log.Panic("Could not get accommodation with id: ", updatePriceDTO.AccommodationId)
		return false, err
	}

	price := models.Price{Amount: updatePriceDTO.NewPrice, Duration: updatePriceDTO.NewInterval}

	reservations, _ := grpcclient.GetReservationsByAccommodationId(updatePriceDTO.AccommodationId)
	for _, reservation := range reservations {
		reservation_interval := models.TimeInterval{Start: reservation.StartDate.AsTime(), End: reservation.EndDate.AsTime()}
		if price.Duration.OtherIntervalOverlaps(reservation_interval) {
			return false, nil
		}
	}

	accommodation.UnitPrice = price
	if err := UpdateAccommodation(*accommodation); err != nil {
		log.Panic("Could not update accommodation with new price. Error: ", err)
		return false, err
	}

	return true, nil
}

func UpdateAvailabilityInterval(updateAvailabilityDTO dtos.UpdateAvailabilityDTO) (bool, error) {
	accommodation, err := repos.GetAccommodationById(updateAvailabilityDTO.AccommodationId)
	if err != nil {
		log.Panic("Could not get accommodation with id: ", updateAvailabilityDTO.AccommodationId)
		return false, err
	}

	reservations, _ := grpcclient.GetReservationsByAccommodationId(updateAvailabilityDTO.AccommodationId)
	for _, reservation := range reservations {
		reservation_interval := models.TimeInterval{Start: reservation.StartDate.AsTime(), End: reservation.EndDate.AsTime()}
		if updateAvailabilityDTO.NewInterval.OtherIntervalOverlaps(reservation_interval) {
			return false, nil
		}
	}

	accommodation.Availability = updateAvailabilityDTO.NewInterval
	if err := UpdateAccommodation(*accommodation); err != nil {
		log.Panic("Could not update accommodation with new availability. Error: ", err)
		return false, err
	}

	return true, nil
}
