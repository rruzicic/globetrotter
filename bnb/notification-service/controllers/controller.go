package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/notification-service/socket"
)

func GetReservationById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}
	//TODO: service part with the fetch logic

	ctx.JSON(200, "Success")
}

func HealthCheck(ctx *gin.Context) {
	socket.SendNotification("This is the title", "This is the message", "user@email.com")
	ctx.JSON(400, "Healthy")
}