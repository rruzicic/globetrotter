package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func RegisterUser(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}
	if services.RegisterUser(user) {
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

func AddUserAPIKey(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}
	addUserAPIKeyDTO := dto.AddUserApiKeyDTO{}

	if err := httpGin.Context.ShouldBindJSON(&addUserAPIKeyDTO); err != nil {
		httpGin.BadRequest(err.Error())
		return
	}

	user, err := services.FindUserByMail(addUserAPIKeyDTO.UserMail)

	if err != nil {
		httpGin.BadRequest(nil)
		return
	}

	if !services.AddUserAPIKey(*user, addUserAPIKeyDTO.APIKey) {
		// should be Internal Server Error but can't see it for some reason
		httpGin.BadRequest(nil)
		return
	}

	httpGin.OK("API Key Added")
}
