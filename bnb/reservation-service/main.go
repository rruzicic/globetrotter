package main

import (
	"github.com/gin-gonic/gin"
	grpcserver "github.com/rruzicic/globetrotter/bnb/accommodation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
)

func main() {
	repos.Connect()
	grpcserver.InitServer()
	ginSetup()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()

	r.Group("/reservation")
	r.POST("/", controllers.CreateReservation)
	r.GET("/:id", controllers.GetReservationById)
	r.GET("/user/:id", controllers.GetReservationsByUserId)
	r.DELETE("/:id", controllers.DeleteReservation)

	r.Run(":8080")
}
