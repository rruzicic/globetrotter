package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/dto"
	"github.com/rruzicic/globetrotter/bnb/account-service/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, "Healthy")
}

func RegisterHost(ctx *gin.Context) {
	var userDto dto.RegisterUserDTO
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(400, "Could not unmarshal request body, error: "+err.Error())
		return
	}
	createdUser, err := services.RegisterHost(dto.RegisterUserDTOtoUser(userDto))
	if err != nil {
		ctx.JSON(400, "Could not register user, error:"+err.Error())
		return
	}
	ctx.JSON(200, createdUser)
}

func RegisterGuest(ctx *gin.Context) {
	var userDto dto.RegisterUserDTO
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(400, "Could not unmarshal request body, error: "+err.Error())
		return
	}
	createdUser, err := services.RegisterGuest(dto.RegisterUserDTOtoUser(userDto))
	if err != nil {
		ctx.JSON(400, "Could not register user, error: "+err.Error())
		return
	}
	ctx.JSON(200, createdUser)
}

func Login(ctx *gin.Context) {
	credentials := dto.CredentialsDTO{}
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(400, "Could not unmarshal request body, error: "+err.Error())
		return
	}
	token, err := services.LoginUser(credentials)
	if err != nil {
		ctx.JSON(400, "Could not login user, error: "+err.Error())
		return
	}
	ctx.JSON(200, token)
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, services.GetAll())
}

func GetById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "Could not convert to ObjectId, error: "+err.Error())
		return
	}
	user, err := services.GetById(id)
	if err != nil {
		ctx.JSON(400, "Could not get user by id, error: "+err.Error())
		return
	}
	ctx.JSON(200, user)
}

func GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := services.GetByEmail(email)
	if err != nil {
		ctx.JSON(400, "Could not get user by email, error: "+err.Error())
		return
	}
	ctx.JSON(200, user)
}

func UpdateUser(ctx *gin.Context) {
	var userDto dto.UpdateUserDTO
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(400, "Could not unmarshal request body, error: "+err.Error())
		return
	}
	updatedUser, err := services.UpdateUser(dto.UpdateUserDTOtoUser(userDto))
	if err != nil {
		ctx.JSON(400, "Could not update user, error: "+err.Error())
		return
	}
	ctx.JSON(200, updatedUser)
}

func DeleteUser(ctx *gin.Context) {
	// TODO: implement
}
