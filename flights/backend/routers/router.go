package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user/register", controllers.RegisterUser)
	r.GET("/hello", controllers.Hello)
	r.GET("/user/all", controllers.GetAllUsers)

	r.Run()
	return nil
}
