package services

import (
	"log"

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

func UpdateAvailabilityInterval(accommodation_id string, new_price float32, new_interval models.TimeInterval) (bool, error) {
	accommodation, err := repos.GetAccommodationById(accommodation_id)
	if err != nil {
		log.Panic("Could not get accommodation with id: ", accommodation_id)
		return false, err
	}

	price := models.Price{Amount: new_price, Duration: new_interval}

	// connect to reservations grpc server
	// for _, res_id := range accommodation.Reservations
	// reservation = reservationsClient.getReservationById
	// if price.interval.otherintervaloverlaps return false, nil
	// disconnect from grpc server
	// accommodation.unitprice = price
	// err := updateaccommodation(accommodation); if err != nil
	// log.panic(nesto)
	// return false, err

	return true, nil
}

func UpdatePriceInterval(accommodation_id string, new_interval models.TimeInterval) (bool, error) {
	accommodation, err := repos.GetAccommodationById(accommodation_id)
	if err != nil {
		log.Panic("Could not get accommodation with id: ", accommodation_id)
		return false, err
	}

	// connect to reservations grpc server
	// for _, res_id := range accommodation.Reservations
	// reservation = reservationsClient.getReservationById
	// if new_interval.otherintervaloverlaps return false, nil
	// disconnect from grpc server
	// accommodation.availability = new_interval
	// err := updateaccommodation(accommodation); if err != nil
	// log.panic(nesto)
	// return false, err

	return true, nil
}
