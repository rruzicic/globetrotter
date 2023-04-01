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

	public.GET("/hello", controllers.Hello)

	public.POST("/user/register", controllers.RegisterUser)
	public.POST("/user/login", controllers.Login)
	adminProtected.GET("/user/all", controllers.GetAllUsers)
	userProtected.GET("/user/current", controllers.CurrentUser) // change possibly to get both user and admin auth
	userProtected.POST("/user/add-api-key-to-user", controllers.AddUserAPIKey)

	public.GET("/flights", controllers.GetAllFlights)
	public.GET("/flights/search", controllers.SearchFlights)
	public.POST("/flights/get-one", controllers.GetFlightById)
	adminProtected.POST("/flights/create", controllers.CreateFlight)
	adminProtected.DELETE("/flights/delete", controllers.DeleteFlight)
	userProtected.POST("/flights/buy-ticket", controllers.BuyTicket)
	userProtected.GET("/flights/get-tickets-by-user", controllers.GetTicketsByUser)

	useAPIKeyMiddleware.POST("/flights/buy-ticket-for-other-user", controllers.BuyTicketForOtherUser)

	userProtected.GET("/api-key", controllers.CreateAPIKey)

	r.Run(config.Configuration.GetString("PORT"))
	return nil
}
