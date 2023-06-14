package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/dtos"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
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

func GetRecommendedAccommodations(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Println("Can't unmarshall given json. Error: ", err.Error())
		ctx.JSON(400, "Bad Request")
		return
	}

	accommodations, err := services.GetRecommendedAccommodations(user)
	if err != nil {
		log.Print("Error in getting recommended accommodations. Error: ", err.Error())
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodations)
	return
}
