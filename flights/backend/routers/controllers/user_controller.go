package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println("Bind err")
	}
	services.RegisterUser(user)
	ctx.JSON(http.StatusCreated, http.Response{
		Status: "203",
	})
}

func GetAllUsers(ctx *gin.Context) {
	users := services.GetAllUsers()
	ctx.JSON(200, users)
}

func Hello(ctx *gin.Context) {
	ctx.Writer.WriteString("Hello!")
}
