package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/config"
	"github.com/rruzicic/globetrotter/flights/backend/controllers"
	"github.com/rruzicic/globetrotter/flights/backend/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())

	public := r.Group("")

	userProtected := r.Group("")
	userProtected.Use(middlewares.UserAuthMiddleware())

	adminProtected := r.Group("")
	adminProtected.Use(middlewares.AdminAuthMiddleware())

	useAPIKeyMiddleware := r.Group("")
	useAPIKeyMiddleware.Use(middlewares.APIKeyAuthMiddleware())

	r.POST("/user/add-api-key-to-user", controllers.AddUserAPIKey)

	public.POST("/user/register", controllers.RegisterUser)
	public.POST("/user/login", controllers.Login)
	public.GET("/hello", controllers.Hello)
	public.GET("/user/all", controllers.GetAllUsers)

	public.GET("/flights/search", controllers.SearchFlights)
	userProtected.GET("/user/current", controllers.CurrentUser)

	public.POST("/flights/create", controllers.CreateFlight)
	public.DELETE("/flights/delete", controllers.DeleteFlight)
	public.GET("/flights", controllers.GetAllFlights)
	public.POST("/flights/get-one", controllers.GetFlightById)

	public.POST("/flights/buy-ticket", controllers.BuyTicket)
	public.GET("/flights/get-tickets-by-user", controllers.GetTicketsByUser)
	useAPIKeyMiddleware.POST("/flights/buy-ticket-for-other-user", controllers.BuyTicketForOtherUser)

	r.GET("/api-key/", controllers.CreateAPIKey)

	r.Run(config.Configuration.GetString("PORT"))
	return nil
}
