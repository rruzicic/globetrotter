package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/dtos"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/services"
)

func CreateHostReview(ctx *gin.Context) {
	var hostReview dtos.CreateHostReviewDTO
	if err := ctx.ShouldBindJSON(&hostReview); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	if _, err := services.CreateHostReview(hostReview); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(201, "Host Review Created")
}

func GetHostReviewById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	hostReview, err := services.GetHostReviewById(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, hostReview)
}

func GetHostReviewsByUserId(ctx *gin.Context){
	userId := ctx.Param("user_id")
	if userId == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	hostReviews, err := services.GetHostReviewsByUserId(userId)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, hostReviews)
}

func GetHostReviewsByHostId(ctx *gin.Context) {
	hostId := ctx.Param("host_id")
	if hostId == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	hostReviews, err := services.GetHostReviewsByHostId(hostId)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, hostReviews)
}

func DeleteHostReview(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	err := services.DeleteHostReview(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Host Review with id: " + id + " successfully deleted.")
}

func UpdateHostReview(ctx *gin.Context){
	var hostReview dtos.CreateHostReviewDTO
	if err := ctx.ShouldBindJSON(&hostReview); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	err := services.UpdateHostReview(hostReview)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Review for host with id: " + hostReview.HostId + " successfully updated.")
}

//==================================================================================================
//==================================================================================================

func CreateAccommodationReview(ctx *gin.Context) {
	var accommodationReview dtos.CreateAccommodationReviewDTO
	if err := ctx.ShouldBindJSON(&accommodationReview); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	if _, err := services.CreateAccommodationReview(accommodationReview); err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(201, "Accommodation Review Created")
}

func GetAccommodationReviewById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	accommodationReview, err := services.GetAccommodationtReviewById(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodationReview)
}

func GetAccommodationReviewsByUserId(ctx *gin.Context){
	userId := ctx.Param("user_id")
	if userId == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	accommodationReviews, err := services.GetAccommodationReviewsByUserId(userId)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodationReviews)
}

func GetAcommodationReviewsByAccommodationId(ctx *gin.Context) {
	accommodationtId := ctx.Param("accommodation_id")
	if accommodationtId == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	accommodationReviews, err := services.GetAccommodationReviewsByAccommodationId(accommodationtId)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, accommodationReviews)
}

func DeleteAccommodationReview(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		ctx.JSON(400, "Bad Request")
		return
	}

	err := services.DeleteAccommodationReview(id)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Accommodation Review with id: " + id + " successfully deleted.")
}

func UpdateAccommodationReview(ctx *gin.Context){
	var accommodationReview dtos.CreateAccommodationReviewDTO
	if err := ctx.ShouldBindJSON(&accommodationReview); err != nil {
		ctx.JSON(400, "Bad Request")
		return
	}

	err := services.UpdateAccommodationReview(accommodationReview)
	if err != nil {
		ctx.JSON(500, "Server Error")
		return
	}

	ctx.JSON(200, "Review for accommodation with id: " + accommodationReview.AccommodationId + " successfully updated.")
}