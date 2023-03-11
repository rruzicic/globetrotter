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
	r.GET("/users/all", controllers.GetAllUsers)

	r.POST("/flights/create", controllers.CreateFlight)
	r.DELETE("/flights/delete", controllers.DeleteFlight)
	r.GET("/flights/", controllers.GetAllFlights)
	r.POST("/flights/get-one", controllers.GetFlightById)

	r.POST("/flights/buy-ticket", controllers.BuyTicket)
	r.GET("/flights/get-tickets-by-user", controllers.GetTicketsByUser)

	r.Run()
	return nil
}
