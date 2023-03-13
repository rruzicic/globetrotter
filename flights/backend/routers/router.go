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

	userProtected := r.Group("")
	userProtected.Use(middlewares.UserAuthMiddleware())

	adminProtected := r.Group("")
	adminProtected.Use(middlewares.AdminAuthMiddleware())

	public := r.Group("")

	public.POST("/user/register", controllers.RegisterUser)
	public.GET("/hello", controllers.Hello)
	public.GET("/user/all", controllers.GetAllUsers)
	public.POST("/user/login", controllers.Login)

	public.POST("/flights/create", controllers.CreateFlight)
	public.DELETE("/flights/delete", controllers.DeleteFlight)
	public.GET("/flights/", controllers.GetAllFlights)
	public.POST("/flights/get-one", controllers.GetFlightById)

	public.POST("/flights/buy-ticket", controllers.BuyTicket)
	public.GET("/flights/get-tickets-by-user", controllers.GetTicketsByUser)

	r.Run()
	return nil
}
