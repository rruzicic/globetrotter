package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/repos"
)

func main() {
	repos.Connect()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()

	r.Group("/accommodation")
	r.POST("/", controllers.CreateAccommodation)
	r.PUT("/", controllers.UpdateAccommodation)

	r.Run(":8080")
}
