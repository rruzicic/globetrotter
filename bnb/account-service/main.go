package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/account-service/gapi"
	"github.com/rruzicic/globetrotter/bnb/account-service/jwt"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
)

func main() {
	repos.Connect()
	go ginSetup()
	gapi.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()

	public := r.Group("/user")

	protected := r.Group("/user")
	protected.Use(jwt.AnyUserAuthMiddleware())

	public.GET("/health", controllers.HealthCheck)
	public.GET("/all", controllers.GetAll)
	public.GET("/id/:id", controllers.GetById)
	public.GET("/email/:email", controllers.GetByEmail)

	public.POST("/register/host", controllers.RegisterHost)
	public.POST("/register/guest", controllers.RegisterGuest)
	protected.POST("/update", controllers.UpdateUser)
	public.POST("/login", controllers.Login)

	protected.DELETE("/delete/:id", controllers.DeleteUser)
	r.Run(":8080")
	log.Println("HTTP server running on port 8080")
}
