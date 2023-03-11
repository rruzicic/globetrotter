package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/DTO"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

type UserIdStruct struct {
	UserId string `json:"userId" bson:"user_id"`
}

func CreateFlight(ctx *gin.Context) {
	var flight models.Flight

	if err := ctx.BindJSON(&flight); err != nil {
		fmt.Println("Passed JSON couldn't be decoded")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusBadRequest, http.Response{
			Status: "400",
		})
	}

	if err := services.CreateFlight(flight); err != nil {
		fmt.Println("Could not save flight document into database")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusInternalServerError, http.Response{
			Status: "500",
		})
	}

	ctx.JSON(http.StatusCreated, http.Response{
		Status: "203",
	})
}

func DeleteFlight(ctx *gin.Context) {
	var flight models.Flight

	if err := ctx.BindJSON(&flight); err != nil {
		fmt.Println("Passed JSON couldn't be decoded")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusBadRequest, http.Response{
			Status: "400",
		})
	}

	if err := services.DeleteFlight(flight); err != nil {
		fmt.Println("Could not delete flight document from database")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusInternalServerError, http.Response{
			Status: "500",
		})
	}

	ctx.JSON(http.StatusOK, http.Response{
		Status: "200",
	})
}

func GetAllFlights(ctx *gin.Context) {
	flights, err := services.GetAllFlights()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.Response{
			Status: "500",
		})
	}

	ctx.JSON(http.StatusOK, flights)
}

func GetFlightById(ctx *gin.Context) {
	var id string

	if err := ctx.BindJSON(id); err != nil {
		fmt.Println("Passed JSON couldn't be decoded")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusBadRequest, http.Response{
			Status: "400",
		})
	}

	flight, err := services.GetFlightById(id)

	if err != nil {
		fmt.Println("Couldn't find flight with id", id)
		fmt.Println(err.Error())

		ctx.JSON(http.StatusInternalServerError, http.Response{
			Status: "500",
		})
	}

	ctx.JSON(http.StatusOK, flight)
}

func BuyTicket(ctx *gin.Context) {
	var request DTO.TicketRequest
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Passed JSON couldn't be decoded")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusBadRequest, http.Response{
			Status: "400",
		})
	}

	err := services.BuyTicket(request.FlightId, request.UserId, request.NumOfTicketsOptional...)

	if err != nil {
		fmt.Println("Couldn't buy ticket")
		fmt.Println(err.Error())

		ctx.JSON(http.StatusInternalServerError, http.Response{
			Status: "500",
		})
	}

	ctx.JSON(http.StatusOK, request)
}

func GetTicketsByUser(ctx *gin.Context) {
	var userIdStruct UserIdStruct
	if err := ctx.BindJSON(&userIdStruct); err != nil {
		log.Println("Passed JSON couldn't be decoded")
		log.Println(err.Error())

		ctx.JSON(http.StatusBadRequest, http.Response{
			Status: "400",
		})
	}

	tickets, err := services.GetTicketsByUser(userIdStruct.UserId)
	if err != nil {
		log.Println("Couldn't get tickets for user")
		log.Println(err.Error())

		ctx.JSON(http.StatusInternalServerError, http.Response{
			Status: "500",
		})
	}

	ctx.JSON(http.StatusOK, tickets)
}
