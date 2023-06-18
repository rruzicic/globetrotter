package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rruzicic/globetrotter/bnb/recommendation-service/dtos"
)

type flightsGinResponse struct {
	Code      int
	Msg       string
	Timestamp uint
	Data      []dtos.Flight
}

func createArrivalUrl(reservationDTO dtos.ReservationDTO, baseUrl string, resource string) string {
	params := url.Values{}
	params.Add("destination", reservationDTO.ArrivalLocationAtReservation)
	params.Add("departure", reservationDTO.DepartureLocationToReservation)
	params.Add("passengerNumber", strconv.Itoa(reservationDTO.People))
	params.Add("departureDateTime", reservationDTO.ReservationStartDate.Format(time.RFC3339))
	params.Add("arrivalDateTime", "")
	url, _ := url.ParseRequestURI(baseUrl)
	url.Path = resource
	url.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", url)

	return urlStr
}

func createDepartureUrl(reservationDTO dtos.ReservationDTO, baseUrl string, resource string) string {
	params := url.Values{}
	params.Add("destination", reservationDTO.ArrivalLocationAtHome)
	params.Add("departure", reservationDTO.DepartureLocationFromReservation)
	params.Add("passengerNumber", strconv.Itoa(reservationDTO.People))
	params.Add("departureDateTime", reservationDTO.ReservationEndDate.Format(time.RFC3339))
	params.Add("arrivalDateTime", "")
	url, _ := url.ParseRequestURI(baseUrl)
	url.Path = resource
	url.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", url)

	return urlStr
}

func getArrivalFlights(reservationDTO dtos.ReservationDTO, baseUrl string, resource string) ([]dtos.Flight, error) {
	url := createArrivalUrl(reservationDTO, baseUrl, resource)

	res, err := http.Get(url)
	if err != nil {
		log.Print("Error sending get request. Error: ", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print("Error reading arrival response's body. Error: ", err.Error())
		return nil, err
	}

	var ginResponse flightsGinResponse
	if err := json.Unmarshal(body, &ginResponse); err != nil {
		log.Print("Could not unmarshall response body into response structure. Error: ", err.Error())
	}

	return ginResponse.Data, nil
}

func getDepartureFlights(reservationDTO dtos.ReservationDTO, baseUrl string, resource string) ([]dtos.Flight, error) {
	url := createDepartureUrl(reservationDTO, baseUrl, resource)

	res, err := http.Get(url)
	if err != nil {
		log.Print("Error sending get request. Error: ", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print("Error reading departure response's body. Error: ", err.Error())
		return nil, err
	}

	var ginResponse flightsGinResponse
	if err := json.Unmarshal(body, &ginResponse); err != nil {
		log.Print("Could not unmarshall response body into response structure. Error: ", err.Error())
	}

	return ginResponse.Data, nil
}

func SearchFlights(reservationDTO dtos.ReservationDTO) ([]dtos.Flight, error) {
	baseURL := "http://flights-backend:8080/flights"
	resource := "/flights/search"

	arrivalFlights, err := getArrivalFlights(reservationDTO, baseURL, resource)
	if err != nil {
		log.Print("Error with getting arrival flights")
		return nil, err
	}

	departureFlights, err := getDepartureFlights(reservationDTO, baseURL, resource)
	if err != nil {
		log.Print("Error with getting departure flights")
		return nil, err
	}

	return append(arrivalFlights, departureFlights...), nil
}

func BuyTicketFromBnBDTO(buyFromBnBDTO dtos.BuyTicketFromBnBDTO) error {
	baseURL := "http://flights-backend:8080/flights"
	resource := "/flights/buy-ticket-for-friend"
	url, _ := url.ParseRequestURI(baseURL)
	url.Path = resource
	urlStr := fmt.Sprintf("%v", url)

	reqBody, err := json.Marshal(buyFromBnBDTO)
	if err != nil {
		log.Print("Could not marshal buy from bnb dto. Error: ", err.Error())
	}

	res, err := http.Post(urlStr, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Print("Could not send request. Error: ", err.Error())
		return err
	}

	return res.Request.Context().Err()
}
