package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/dtos"
)

func SearchFlights(ctx *gin.Context) {
	var reservationDTO dtos.ReservationDTO
	if err := ctx.ShouldBindJSON(&reservationDTO); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}
}
