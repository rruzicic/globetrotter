package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func RegisterUser(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	user := dto.RegisterUserDTO{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}
	if services.RegisterUser(dto.RegisterUserDTOToUser(user)) {
		httpGin.Created(nil)
		return
	}
	httpGin.BadRequest(nil)
}

func GetAllUsers(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	httpGin.OK(services.GetAllUsers())
}

func Hello(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	httpGin.OK("Hello")
}
