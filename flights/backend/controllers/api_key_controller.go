package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func GenerateAPIKey(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}

	temporary_str := httpGin.Context.Query("temporary")
	temporary_bool, err := strconv.ParseBool(temporary_str)
	if err != nil {
		httpGin.BadRequest("Value of query \"temporary\" not a boolean string")
		return
	}

	key := services.GenerateAPIKey(temporary_bool)

	httpGin.OK(key)
	return
}

func APIKeyExpired(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}

	var key models.APIKey
	if err := ctx.ShouldBindJSON(&key); err != nil {
		httpGin.BadRequest("Could not bind json to api key")
		return
	}

	expired := services.APIKeyExpired(key)

	httpGin.OK(expired)
	return
}

func AddAPIKeyToUser(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}

	var key models.APIKey
	if err := ctx.ShouldBindJSON(&key); err != nil {
		httpGin.BadRequest("Could not bind json to add key DTO")
		return
	}

	user, err := services.GetUserFromToken(ctx)
	if err != nil {
		httpGin.Unauthorized("User token not valid")
		return
	}

	success := services.AddAPIKeyToUser(*user, key)
	if success {
		httpGin.OK("Key added to user")
		return
	}
	httpGin.InternalServerError("API Key not added to user")
	return
}

func FindUserByAPIKey(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}

	var key models.APIKey
	if err := ctx.ShouldBindJSON(&key); err != nil {
		httpGin.BadRequest("Could not bind api key from json")
		return
	}

	user, err := services.FindUserByAPIKey(key)
	if err != nil {
		httpGin.InternalServerError("Could not find user")
		return
	}

	httpGin.OK(user)
	return
}

func BuyTicketForFriend(ctx *gin.Context) {
	httpGin := http.Gin{Context: ctx}

	var buyTicketDTO dto.BuyTicketForOtherUserDTO
	if err := ctx.ShouldBindJSON(&buyTicketDTO); err != nil {
		httpGin.BadRequest("Could not bind json to buy ticket for friend DTO")
		return
	}

	if err := services.BuyTicketForFriend(buyTicketDTO.FlightId, buyTicketDTO.ApiKey, buyTicketDTO.NumOfTicketsOptional...); err != nil {
		httpGin.InternalServerError("Could not buy ticket for friend")
		return
	}

	httpGin.OK("Ticket bought")
	return
}
