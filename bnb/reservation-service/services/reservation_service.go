package services

import (
	"log"
	"time"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/dtos"
	grpcclient "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReservation(reservationDTO dtos.CreateReservationDTO) (bool, error) {
	acc_id, err := primitive.ObjectIDFromHex(reservationDTO.AccommodationId)
	if err != nil {
		return false, err
	}

	user_id, err := primitive.ObjectIDFromHex(reservationDTO.UserId)
	if err != nil {
		return false, err
	}

	reservation := models.Reservation{
		AccommodationId: &acc_id,
		UserId:          &user_id,
		DateInterval:    reservationDTO.DateInterval,
		NumOfGuests:     reservationDTO.NumOfGuests,
		IsApproved:      false,
		TotalPrice:      0.0,
	}

	accommodation, err := grpcclient.GetAccommodationById(reservation.AccommodationId.Hex())
	if err != nil {
		return false, err
	}

	if accommodation.PriceForPerson {
		reservation.TotalPrice = float32(reservation.NumOfGuests) * accommodation.Amount
	} else {
		total_days := reservation.DateInterval.Start.Sub(reservation.DateInterval.End).Hours() / 24
		reservation.TotalPrice = float32(total_days) * accommodation.Amount
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
	return repos.GetActiveReservationsByUser(id)
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
	_, err := grpcclient.IncrementCancellationsCounter(id)
	if err != nil {
		return err
	}
	return repos.DeleteReservation(id)
}

func ApproveReservation(id string) error {
	reservation, err := repos.GetReservationById(id)
	if err != nil {
		log.Panic("Could not get reservation by id. Error: ", err)
		return err
	}

	reservation.IsApproved = true

	// reject others that overlap
	accommodation, err := grpcclient.GetAccommodationById(reservation.AccommodationId.Hex())
	if err != nil {
		log.Panic("Could not get accommodation by id from accommodation service. Error: ", err)
		return err
	}

	for _, res_id := range accommodation.Reservations {
		if res_id != id {
			existing_reservation, err := repos.GetReservationById(res_id)
			if err != nil {
				log.Panic("Could not get reservation by id. Error: ", err)
				return err
			}

			// this makes no sense right now, but if there was an enum this would be like pending or sth
			if existing_reservation.DateInterval.OtherIntervalOverlaps(reservation.DateInterval) && existing_reservation.IsApproved == false {
				RejectReservation(existing_reservation.Id.Hex())
			}
		}
	}

	return repos.UpdateReservation(*reservation)
}

func RejectReservation(id string) error {
	reservation, err := repos.GetReservationById(id)
	if err != nil {
		log.Panic("Could not get reservation by id. Error: ", err)
		return err
	}

	reservation.IsApproved = false
	return repos.UpdateReservation(*reservation)
}

func GetReservationsByAccommodationId(id string) ([]models.Reservation, error) {
	reservations, err := repos.GetReservationsByAccommodationId(id)
	if err != nil {
		log.Panic("Could not get reservations by accommodation id: ", id)
		return nil, err
	}

	return reservations, err
}
