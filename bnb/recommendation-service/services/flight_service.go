package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rruzicic/globetrotter/bnb/recommendation-service/dtos"
)

func SearchFlights(reservationDTO dtos.ReservationDTO) ([]dtos.Flight, error) {
	arrival_query := fmt.Sprintf("destination=%s&departure=%s&passengerNumber=%d&departureDateTime=%s&arrivalDateTime=%s", reservationDTO.ArrivalDestination, "", reservationDTO.People, reservationDTO.ReservationStartDate.String(), "")
	departure_query := fmt.Sprintf("destination=%s&departure=%s&passengerNumber=%d&departureDateTime=%s&arrivalDateTime=%s", reservationDTO.DepartureDestination, "", reservationDTO.People, reservationDTO.ReservationStartDate.String(), "")
	var arrival_flights []dtos.Flight
	var departure_flights []dtos.Flight

	arrival_flights_res, err := http.Get(fmt.Sprintf("http://flights-backend:8080/flights/search?%s", arrival_query))
	if err != nil {
		return nil, err
	}
	defer arrival_flights_res.Body.Close()
	arrival_flights_res_body, _ := ioutil.ReadAll(arrival_flights_res.Body)
	if err := json.Unmarshal(arrival_flights_res_body, &arrival_flights); err != nil {
		return nil, err
	}

	departure_flights_res, err := http.Get(fmt.Sprintf("http://flights-backend:8080/flights/search?%s", departure_query))
	if err != nil {
		return nil, err
	}
	defer departure_flights_res.Body.Close()
	departure_flights_res_body, _ := ioutil.ReadAll(departure_flights_res.Body)
	if err := json.Unmarshal(departure_flights_res_body, &departure_flights); err != nil {
		return nil, err
	}

	return append(arrival_flights, departure_flights...), nil
}
