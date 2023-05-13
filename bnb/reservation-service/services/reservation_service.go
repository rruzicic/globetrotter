package services

import (
	"time"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
)

func CreateReservation(reservation models.Reservation) error {
	return repos.CreateReservation(reservation)
}

func GetReservationById(id string) (*models.Reservation, error) {
	return repos.GetReservationById(id)
}

func GetReservationsByUserId(id string) ([]models.Reservation, error) {
	return repos.GetReservationsByUserId(id)
}

func GetFutureActiveReservationsByHost(id string) ([]models.Reservation, error) {
	reservations, err := GetReservationsByHostId(id)
	if err != nil {
		return []models.Reservation{}, nil
	}
	var futureApprovedReservations []models.Reservation
	for _, reservation := range reservations {
		if reservation.DateInterval.DateIsAfter(time.Now()) && reservation.IsApproved {
			futureApprovedReservations = append(futureApprovedReservations, reservation)
		}
	}
	return futureApprovedReservations, nil
}

func GetActiveReservationsByUser(id string) ([]models.Reservation, error) {
	reservations, err := repos.GetReservationsByUserId(id)
	if err != nil {
		return []models.Reservation{}, err
	}
	var activeReservations []models.Reservation
	for _, reservation := range reservations {
		if reservation.DateInterval.DateIsAfter(time.Now()) && reservation.IsApproved {
			activeReservations = append(activeReservations, reservation)
		}
	}
	return activeReservations, nil
}

func GetReservationsByHostId(id string) ([]models.Reservation, error) {
	// get all accomodations from accomodation service by host id
	accomodations, err := getAccomodationsByHostId(id)
	if err != nil {
		return []models.Reservation{}, err
	}
	var futureApprovedReservations []models.Reservation
	for _, accomodation := range accomodations {
		reservation := repos.GetReservationsByAccommodationId(accomodation.Id)
		if reservation.DateInterval.DateIsAfter(time.Now()) && reservation.IsApproved {
			futureApprovedReservations = append(futureApprovedReservations, reservation)
		}
	}
	return futureApprovedReservations, nil
	// for accomodation in accomodations
	//reservation := repos.GetReservationsByAccommodationId(accomodationId)
	//if reservation.DateInterval.DateIsAfter(time.Now()) && reservation.IsApproved {
	//	futureApprovedReservations = append(futureApprovedReservations, reservation)
	//}

}

func DeleteReservation(id string) error {
	return repos.DeleteReservation(id)
}

func getAccomodationsByHostId(id string) error {
	return nil
}
