package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func CreateAPIKey(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	api_key_value := services.CreateAPIKey()

	httpGin.OK(api_key_value)
}
