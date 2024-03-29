package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/dtos"
	grpcclient "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateReservation(reservationDTO dtos.CreateReservationDTO) (*models.Reservation, error) {
	acc_id, err := primitive.ObjectIDFromHex(reservationDTO.AccommodationId)
	if err != nil {
		return nil, err
	}

	user_id, err := primitive.ObjectIDFromHex(reservationDTO.UserId)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	_, err = grpcclient.ReservationCreated(reservation, accommodation.Name, accommodation.User)
	if err != nil {
		return nil, err
	}

	accommodation_availability := models.TimeInterval{Start: accommodation.AvailabilityStartDate.AsTime(), End: accommodation.AvailabilityEndDate.AsTime()}
	if !accommodation_availability.OtherIntervalIsDuring(reservation.DateInterval) {
		err := errors.New("Reservation date isn't during accommodations' availability")
		log.Print(err.Error())
		return nil, err
	}

	if accommodation.Guests < int32(reservationDTO.NumOfGuests) {
		err := errors.New("Number of guests greater than accommodations' capacity")
		log.Print(err.Error())
		return nil, err
	}

	if accommodation.PriceForPerson {
		reservation.TotalPrice = float32(reservation.NumOfGuests) * accommodation.Amount
	} else {
		total_days := reservation.DateInterval.End.Sub(reservation.DateInterval.Start).Hours() / 24
		reservation.TotalPrice = float32(total_days) * accommodation.Amount
	}

	// check if there are overlapping active reservations
	for _, reservation_id := range accommodation.Reservations {
		existing_reservation, err := repos.GetReservationById(reservation_id)
		if err != nil {
			return nil, err
		}

		if existing_reservation.DateInterval.OtherIntervalOverlaps(reservation.DateInterval) && (existing_reservation.IsApproved == true) {
			err := errors.New("Reservation exists in that time")
			log.Print(err.Error())
			return nil, err
		}
	}

	// Check if autoaccept
	if accommodation.AutoApprove == true {
		reservation.IsApproved = true
	} else {
		reservation.IsApproved = false
	}

	returnValue, err := repos.CreateReservation(reservation)
	if err != nil {
		return nil, err
	}
	//Get Host Id
	hostAnswer, _ := grpcclient.GetHostByAccommodation(reservation.AccommodationId.Hex())

	//Publish an event to the account service
	conn := Conn()
	defer conn.Close()

	event := pb.ReservationEvent{
		AccommodationId: returnValue.AccommodationId.Hex(),
		UserId:          returnValue.UserId.Hex(),
		StartDate:       timestamppb.New(returnValue.DateInterval.Start),
		EndDate:         timestamppb.New(returnValue.DateInterval.End),
		NumOfGuests:     int32(returnValue.NumOfGuests),
		IsApproved:      returnValue.IsApproved,
		TotalPrice:      returnValue.TotalPrice,
		Id:              returnValue.Id.Hex(),
		HostId:          hostAnswer.HostId,
	}
	data, _ := proto.Marshal(&event)
	err = conn.Publish("account-service-2", data)
	if err != nil {
		log.Panic(err)
	}

	return returnValue, nil
}

func GetReservationById(id string) (*models.Reservation, error) {
	return repos.GetReservationById(id)
}

func GetReservationsByUserId(id string) ([]models.Reservation, error) {
	return repos.GetReservationsByUserId(id)
}

/*func GetFutureActiveReservationsByHost(id string) ([]models.Reservation, error) {
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
}*/

func GetActiveReservationsByUser(id string) ([]models.Reservation, error) {
	return repos.GetActiveReservationsByUser(id)
}

func GetFinishedReservationsByUser(id string) ([]models.Reservation, error) {
	return repos.GetFinishedReservationsByUser(id)
}

func GetFutureActiveReservationsByHost(id string) ([]models.Reservation, error) {
	accomodations, err := grpcclient.GetAccommodationByHostId(id)
	//log.Println("accomodations for given host id: ", accomodations)
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
			if reservation.DateInterval.DateIsBefore(time.Now()) && reservation.IsApproved {
				futureApprovedReservations = append(futureApprovedReservations, reservation)
			}
		}
	}
	return futureApprovedReservations, nil
}

func DeleteReservation(id string) error {
	reservation, err := repos.GetReservationById(id)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	if err := repos.DeleteReservation(id); err != nil {
		log.Print(err.Error())
		return err
	}

	hostAnswer, _ := grpcclient.GetHostByAccommodation(reservation.AccommodationId.Hex())

	accommodation, err := grpcclient.GetAccommodationById(reservation.AccommodationId.Hex())
	if err != nil {
		return err
	}

	//Publish an event to the account service
	conn := Conn()
	defer conn.Close()

	event := pb.ReservationEvent{
		AccommodationId:   reservation.AccommodationId.Hex(),
		UserId:            reservation.UserId.Hex(),
		StartDate:         timestamppb.New(reservation.DateInterval.Start),
		EndDate:           timestamppb.New(reservation.DateInterval.End),
		NumOfGuests:       int32(reservation.NumOfGuests),
		IsApproved:        reservation.IsApproved,
		TotalPrice:        reservation.TotalPrice,
		Id:                reservation.Id.Hex(),
		HostId:            hostAnswer.HostId,
		AccommodationName: accommodation.Name,
	}
	data, _ := proto.Marshal(&event)
	_, err = conn.Subscribe("saga-cancel-reservation-2", func(message *nats.Msg) {
		if string(message.Data) == "OK" {
			_, err = conn.Subscribe("saga-cancel-reservation-4", func(message *nats.Msg) {
				if string(message.Data) == "OK" {
					res, err := grpcclient.IncrementCancellationsCounter(reservation.UserId.Hex())
					if err != nil {
						log.Print(res)
					}

					boolAns, err := grpcclient.RemoveReservationFromAccommodation(reservation.AccommodationId.Hex(), id)
					if err != nil {
						log.Print(boolAns, err.Error())
					}
				} else {
					//rollback account-service stuff that happened and reservation canceling
					err = conn.Publish("saga-cancel-rollback-account", data)
					if err != nil {
						log.Panic(err)
					}
					repos.CreateReservation(*reservation)
					return

				}
			})
			err = conn.Publish("saga-cancel-reservation-3", data)
			if err != nil {
				log.Panic(err)
			}
		} else {
			//rollback reservation cancelling, or cancel canceling of reservation
			repos.CreateReservation(*reservation)
			return
		}
	})
	err = conn.Publish("saga-cancel-reservation-1", data)
	if err != nil {
		log.Panic(err)
	}

	// _, err = grpcclient.ReservationCanceled(*reservation, accommodation.Name, accommodation.User)
	// if err != nil {
	// 	log.Print(err.Error())
	// }

	return nil
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

	grpcclient.ReservationResponse(*reservation, accommodation.Name)

	return repos.UpdateReservation(*reservation)
}

func RejectReservation(id string) error {
	reservation, err := repos.GetReservationById(id)
	if err != nil {
		log.Panic("Could not get reservation by id. Error: ", err)
		return err
	}
	accommodation, err := grpcclient.GetAccommodationById(reservation.AccommodationId.Hex())
	if err != nil {
		log.Panic("Could not get accommodation by id from accommodation service. Error: ", err)
		return err
	}

	reservation.IsApproved = false

	grpcclient.ReservationResponse(*reservation, accommodation.Name)

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

func Conn() *nats.Conn {
	conn, err := nats.Connect("nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
