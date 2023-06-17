package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/controllers"
	grpcserver "github.com/rruzicic/globetrotter/bnb/recommendation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/repos"
)

func main() {
	go ginSetup()
	repos.Connect()
	grpcserver.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	rec := r.Group("/recommendation")
	rec.POST("/flights", controllers.SearchFlights)
	rec.POST("/accommodations", controllers.GetRecommendedAccommodations)
	rec.GET("/accommodations/init", controllers.InitDBData)
	rec.GET("/accommodations/drop", controllers.DropDB)
	rec.GET("/accommodations/init-mock", controllers.LoadMockData)

	r.Run(":8080")
}
