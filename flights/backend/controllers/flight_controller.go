package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

type UserIdStruct struct {
	UserId string `json:"userId" bson:"user_id"`
}

func CreateFlight(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	var flight models.Flight

	if err := ctx.BindJSON(&flight); err != nil {
		log.Println("Passed JSON couldn't be decoded")
		log.Println(err.Error())

		httpGin.BadRequest(nil)
		return
	}

	if err := services.CreateFlight(flight); err != nil {
		log.Println("Could not save flight document into database")
		log.Println(err.Error())

		httpGin.NoContent(nil)
		return
	}

	httpGin.Created(nil)
}

func DeleteFlight(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	var flight models.Flight

	if err := ctx.BindJSON(&flight); err != nil {
		log.Println("Passed JSON couldn't be decoded")
		log.Println(err.Error())

		httpGin.BadRequest(nil)
		return
	}

	if err := services.DeleteFlight(flight); err != nil {
		log.Println("Could not delete flight document from database")
		log.Println(err.Error())

		httpGin.NoContent(nil)
		return
	}

	httpGin.OK(nil)
}

func GetAllFlights(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	flights, err := services.GetAllFlights()

	if err != nil {
		httpGin.NoContent(nil)
		return
	}

	httpGin.OK(flights)
}

func GetFlightById(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	var id string

	if err := ctx.BindJSON(id); err != nil {
		log.Println("Passed JSON couldn't be decoded")
		log.Println(err.Error())

		httpGin.BadRequest(nil)
		return
	}

	flight, err := services.GetFlightById(id)

	if err != nil {
		log.Println("Couldn't find flight with id", id)
		log.Println(err.Error())

		httpGin.NoContent(nil)
		return
	}

	httpGin.OK(flight)
}

func BuyTicket(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	var request dto.TicketRequest
	if err := ctx.BindJSON(&request); err != nil {
		log.Println("Passed JSON couldn't be decoded")
		log.Println(err.Error())

		httpGin.BadRequest(nil)
		return
	}

	err := services.BuyTicket(request.FlightId, request.UserId, request.NumOfTicketsOptional...)

	if err != nil {
		log.Println("Couldn't buy ticket")
		log.Println(err.Error())

		httpGin.NoContent(nil)
		return
	}

	httpGin.OK(request)
}

func GetTicketsByUser(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	var userIdStruct UserIdStruct
	if err := ctx.BindJSON(&userIdStruct); err != nil {
		log.Println("Passed JSON couldn't be decoded")
		log.Println(err.Error())

		httpGin.BadRequest(nil)
		return
	}

	tickets, err := services.GetTicketsByUser(userIdStruct.UserId)
	if err != nil {
		log.Println("Couldn't get tickets for user")
		log.Println(err.Error())

		httpGin.NoContent(nil)
		return
	}

	httpGin.OK(tickets)
}
