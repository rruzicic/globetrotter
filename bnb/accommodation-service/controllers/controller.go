package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/dtos"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/services"
)

func CreateAccommodation(ctx *gin.Context) {
	var accommodation dtos.CreateAccommodationDTO
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

func GetAllAccommodations(ctx *gin.Context) {
	accommodations, err := services.GetAllAccommodations()
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodations)
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

func UpdatePriceInterval(ctx *gin.Context) {
	var dto dtos.UpdatePriceDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	retval, err := services.UpdatePriceInterval(dto)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	if retval == true {
		ctx.JSON(200, "Accommodation price updated")
		return
	}

	ctx.JSON(500, "Could not update price")
}

func UpdateAvailabilityInterval(ctx *gin.Context) {
	var dto dtos.UpdateAvailabilityDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	retval, err := services.UpdateAvailabilityInterval(dto)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	if retval == true {
		ctx.JSON(200, "Accommodation availability updated")
		return
	}

	ctx.JSON(500, "Could not update availability")
}

func SearchAccomodation(ctx *gin.Context) {
	cityName := ctx.DefaultQuery("cityName", "")
	guestNum := ctx.DefaultQuery("guestNum", "")
	startDate := ctx.DefaultQuery("startDate", "")
	endDate := ctx.DefaultQuery("endDate", "")
	if startDate == "" || endDate == "" {
		ctx.JSON(400, "Start date and end date do not exist")
		return
	}
	log.Println(cityName, guestNum, startDate, endDate)
	searchResult, err := services.SearchAccomodation(cityName, guestNum, startDate, endDate)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, searchResult)
}

func GetAccommodationsByHostId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	accommodations, err := services.GetAccommodationsByHostId(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodations)
}

func GetAccommodationById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, "Bad Request")
		return
	}

	accommodation, err := services.GetAccommodationById(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodation)
}
