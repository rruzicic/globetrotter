package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/notification-service/repos"
)

func GetNotificationsByUserId(ctx *gin.Context) {
	userId := ctx.Param("id")
	log.Println("User id: ", userId)
	if userId == "" {
		ctx.JSON(400, "Id is required")
		return
	}
	notifications, err := repos.GetNotificationsByUserId(userId)
	log.Println("Num of notifications: ", len(notifications))
	if err != nil {
		log.Panic(err)
	}

	ctx.JSON(200, notifications)
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(400, "Healthy")
}