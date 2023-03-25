package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/controllers"
	"github.com/rruzicic/globetrotter/flights/backend/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	useAPIKeyMiddleware := r.Group("/flights/buy-ticket-for-other-user")
	useAPIKeyMiddleware.Use(middlewares.APIKeyAuthMiddleware())

	r.POST("/user/register", controllers.RegisterUser)
	r.GET("/hello", controllers.Hello)
	r.GET("/user/all", controllers.GetAllUsers)
	r.POST("/user/add-api-key-to-user", controllers.AddUserAPIKey)

	r.POST("/flights/create", controllers.CreateFlight)
	r.DELETE("/flights/delete", controllers.DeleteFlight)
	r.GET("/flights/", controllers.GetAllFlights)
	r.POST("/flights/get-one", controllers.GetFlightById)

	r.POST("/flights/buy-ticket", controllers.BuyTicket)
	r.GET("/flights/get-tickets-by-user", controllers.GetTicketsByUser)
	r.POST("/flights/buy-ticket-for-other-user", controllers.BuyTicketForOtherUser)
	r.GET("/flights/search", controllers.SearchFlights)

	r.GET("/api-key/", controllers.CreateAPIKey)

	r.Run()
	return nil
}
