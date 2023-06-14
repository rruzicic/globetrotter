package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/controllers"
	grpcserver "github.com/rruzicic/globetrotter/bnb/recommendation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/repos"
)

func main() {
	repos.Connect()
	grpcserver.InitServer()
	go ginSetup()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	rec := r.Group("/recommendation", controllers.SearchFlights)
	rec.GET("/flights")
	rec.POST("/accommodations", controllers.GetRecommendedAccommodations)

	r.Run(":8080")
}
