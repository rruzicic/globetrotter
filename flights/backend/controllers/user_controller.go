package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func RegisterUser(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		httpGin.BadRequest()
		log.Print("could not bind JSON to models.User")
	}

	services.RegisterUser(user)
	httpGin.Created()
}

func GetAllUsers(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	httpGin.OKObject(services.GetAllUsers())
}

func Hello(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	httpGin.OKObject("Hello")
}
