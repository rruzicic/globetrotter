package controllers

import "github.com/gin-gonic/gin"

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
	ctx.JSON(400, "Healthy")
}