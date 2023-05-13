package services

import (
	"time"

	grpcclient "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
)

func CreateReservation(reservation models.Reservation) (bool, error) {
	accommodation, err := grpcclient.GetAccommodationById(reservation.AccommodationId.Hex())
	if err != nil {
		return false, err
	}

	// check if there are overlapping active reservations
	for _, reservation_id := range accommodation.Reservations {
		existing_reservation, err := repos.GetReservationById(reservation_id)
		if err != nil {
			return false, err
		}

		if existing_reservation.DateInterval.OtherIntervalOverlaps(reservation.DateInterval) && existing_reservation.IsApproved == true {
			return false, nil
		}
	}

	// Check if autoaccept
	if accommodation.AutoApprove == true {
		reservation.IsApproved = true
	} else {
		reservation.IsApproved = false
	}

	return true, repos.CreateReservation(reservation)
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
	accomodations, err := grpcclient.GetAccommodationByHostId(id)
	if err != nil {
		return []models.Reservation{}, err
	}
	var futureApprovedReservations []models.Reservation
	for _, accomodation := range accomodations {
		reservations, err := repos.GetReservationsByAccommodationId(accomodation.Id)
		if err != nil {
			return []models.Reservation{}, err
		}
		for _, reservation := range reservations {
			if reservation.DateInterval.DateIsAfter(time.Now()) && reservation.IsApproved {
				futureApprovedReservations = append(futureApprovedReservations, reservation)
			}
		}
	}
	return futureApprovedReservations, nil
}

func DeleteReservation(id string) error {
	//TODO: Povecati brojac otkazanih rezervacija u useru. grpc metoda prema Ratku

	return repos.DeleteReservation(id)
}
