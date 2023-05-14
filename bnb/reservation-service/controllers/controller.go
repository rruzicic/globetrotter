package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/dtos"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/services"
)

func CreateReservation(ctx *gin.Context) {
	var reservationDTO dtos.CreateReservationDTO
	if err := ctx.ShouldBindJSON(&reservationDTO); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	retval, err := services.CreateReservation(reservationDTO)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	if retval == true {
		ctx.JSON(201, "Reservation Created")
		return
	}

	ctx.JSON(500, "Server Error")
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

func ApproveReservation(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	if err := services.ApproveReservation(id); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Reservation Approved")
}

func RejectReservation(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	if err := services.RejectReservation(id); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Reservation Rejected")
}
