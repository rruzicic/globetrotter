package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/dto"
	grpcclient "github.com/rruzicic/globetrotter/bnb/account-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/account-service/jwt"
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
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "Could not convert to ObjectId, error: "+err.Error())
		return
	}
	err = services.DeleteUser(id)
	if err != nil {
		ctx.JSON(400, "Could not delete user, error: "+err.Error())
		return
	}
	ctx.JSON(200, "User deleted successfully")
}

func Ping(ctx *gin.Context) {
	msg, err := grpcclient.Ping()
	if err != nil {
		ctx.JSON(400, "Something went wrong")
		return
	}
	ctx.JSON(200, msg)
}

func AddAPIKeyToUser(ctx *gin.Context) {
	key := ctx.Query("key")
	email, err := jwt.ExtractTokenEmail(ctx)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	user, err := services.AddAPIKeyToUser(email, key)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, user)
}
