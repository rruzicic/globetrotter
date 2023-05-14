package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/dtos"
	grpcclient "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/services"
)

func CreateReservation(ctx *gin.Context) {
	var reservationDTO dtos.CreateReservationDTO
	if err := ctx.ShouldBindJSON(&reservationDTO); err != nil {
		log.Print(err.Error())
		ctx.JSON(400, "Bad Request")
		return
	}

	reservation, err := services.CreateReservation(reservationDTO)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(201, reservation)
}

func GetReservationById(ctx *gin.Context) {
	id := ctx.Param("id")
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
	id := ctx.Param("id")
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
	id := ctx.Param("id")
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
	id := ctx.Param("id")
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
	id := ctx.Param("id")
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

func GetReservationsByAccommodationId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	reservations, err := services.GetReservationsByAccommodationId(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, reservations)
}

func TestConnection(ctx *gin.Context) {
	id := ctx.Param("msg")
	log.Print(id)
	grpcclient.TestConnection(id)
}

func AddReservationToAccommodation(ctx *gin.Context) {
	accommodation_id := ctx.Param("acc_id")
	if accommodation_id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	reservation_id := ctx.Param("res_id")
	if reservation_id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	boolAns, err := grpcclient.AddReservationToAccommodation(accommodation_id, reservation_id)
	if err != nil {
		log.Panic("Error adding reservation to accommodation. Error: ", err.Error())
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, boolAns)
}
