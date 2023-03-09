package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/rruzicic/globetrotter/flights/backend/controllers"
	routers "github.com/rruzicic/globetrotter/flights/backend/routers/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user/register", routers.RegisterUser)
	r.GET("/hello", routers.Hello)
	r.GET("/users/all", routers.GetAllUsers)

	r.POST("/flights/create", controllers.CreateFlight)
	r.DELETE("/flights/delete", controllers.DeleteFlight)
	r.GET("/flights/", controllers.GetAllFlights)
	r.POST("/flights/get-one", controllers.GetFlightById)

	r.Run()
	return nil
}
