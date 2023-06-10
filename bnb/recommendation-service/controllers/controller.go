package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/dtos"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/services"
)

func SearchFlights(ctx *gin.Context) {
	var reservationDTO dtos.ReservationDTO
	if err := ctx.ShouldBindJSON(&reservationDTO); err != nil {
		log.Printf("Can't unmarshall given json. Error: %s", err.Error())
		ctx.JSON(400, "Bad Request")
		return
	}

	flights, err := services.SearchFlights(reservationDTO)
	if err != nil {
		log.Printf("Error in searching for flighs. Error: %s", err.Error())
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, flights)
	return
}
