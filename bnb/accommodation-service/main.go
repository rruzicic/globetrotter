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
	go ginSetup()
	grpc_server.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	acc := r.Group("/accommodation")
	acc.POST("/", controllers.CreateAccommodation)
	acc.GET("/", controllers.GetAllAccommodations)
	acc.PUT("/", controllers.UpdateAccommodation)
	acc.PUT("/price", controllers.UpdatePriceInterval)
	acc.PUT("/availability", controllers.UpdateAvailabilityInterval)
	acc.GET("/search", controllers.SearchAccomodation)
	acc.GET("/host/:id", controllers.GetAccommodationsByHostId)
	acc.GET("/:id", controllers.GetAccommodationById)

	r.Run(":8080")
}
