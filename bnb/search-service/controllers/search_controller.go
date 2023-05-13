package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/search-service/services"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, "We good");
}

func Search() {
	services.SearchAccommodation();
}