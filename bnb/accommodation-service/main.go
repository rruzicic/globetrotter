package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/controllers"
	grpc_server "github.com/rruzicic/globetrotter/bnb/accommodation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/repos"
)

func main() {
	repos.Connect()
	grpc_server.InitServer()
	ginSetup()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	r.Group("/accommodation")
	r.POST("/", controllers.CreateAccommodation)
	r.GET("/", controllers.GetAllAccommodations)
	r.PUT("/", controllers.UpdateAccommodation)
	r.PUT("/price", controllers.UpdatePriceInterval)
	r.PUT("/availability", controllers.UpdateAvailabilityInterval)
	r.GET("/search", controllers.SearchAccomodation)
	r.GET("/host/:id", controllers.GetAccommodationsByHostId)

	r.Run(":8080")
}
