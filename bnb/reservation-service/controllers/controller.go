package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/services"
)

func CreateReservation(ctx *gin.Context) {
	var reservation models.Reservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	if err := services.CreateReservation(reservation); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(201, "Reservation Created")
}

func GetReservationById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	reservation, err := services.GetReservationById(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, reservation)
}

func GetReservationsByUserId(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	reservations, err := services.GetReservationsByUserId(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, reservations)
}

func DeleteReservation(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	if err := services.DeleteReservation(id); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Reservation Deleted")
}
