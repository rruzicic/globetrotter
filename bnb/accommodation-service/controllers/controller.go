package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/services"
)

func CreateAccommodation(ctx *gin.Context) {
	var accommodation models.Accommodation
	if err := ctx.ShouldBindJSON(&accommodation); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	if err := services.CreateAccommodation(accommodation); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(201, "Accommodation Created")
}

func UpdateAccommodation(ctx *gin.Context) {
	var accommodation models.Accommodation
	if err := ctx.ShouldBindJSON(&accommodation); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	if err := services.UpdateAccommodation(accommodation); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Accommodation Updated")
}
