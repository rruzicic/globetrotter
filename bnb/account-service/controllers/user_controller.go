package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/models"
	"github.com/rruzicic/globetrotter/bnb/account-service/services"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, "Healthy")
}

func RegisterHost(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, "Bad request")
		return
	}
	ctx.JSON(200, services.RegisterHost(user))
}

func RegisterGuest(ctx *gin.Context) {
	// TODO: implement
}

func Login(ctx *gin.Context) {
	// TODO: implement
}

func GetAll(ctx *gin.Context) {
	// TODO: implement
}

func GetById(ctx *gin.Context) {
	// TODO: implement
}

func GetByEmail(ctx *gin.Context) {
	// TODO: implement
}

func UpdateUser(ctx *gin.Context) {
	// TODO: implement
}

func DeleteUser(ctx *gin.Context) {
	// TODO: implement
}
