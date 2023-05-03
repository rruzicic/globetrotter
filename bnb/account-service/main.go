package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/account-service/gapi"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
)

func main() {
	gapi.InitServer()
	repos.Connect()
	ginSetup()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()

	r.Group("/user")

	r.GET("/health", controllers.HealthCheck)
	r.GET("/all", controllers.GetAll)
	r.GET("/:id", controllers.GetById)
	r.GET("/:email", controllers.GetByEmail)

	r.POST("/register/host", controllers.RegisterHost)
	r.POST("/register/guest", controllers.RegisterGuest)
	r.POST("/update", controllers.UpdateUser)
	r.POST("/login", controllers.Login)

	r.DELETE("/delete/:id", controllers.DeleteUser)
	r.Run(":8080")
}
