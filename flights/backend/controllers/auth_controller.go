package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func Login(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	credentials := dto.LoginDTO{}
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}
	token, err := services.Login(credentials)
	if err != nil {
		httpGin.BadRequest(err.Error())
		return
	}

	httpGin.OK(token)
}
