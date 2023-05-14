package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/controllers"
	grpcserver "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
)

func main() {
	repos.Connect()
	go ginSetup()
	grpcserver.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	r.Group("/reservation")
	r.POST("/", controllers.CreateReservation)
	r.GET("/:id", controllers.GetReservationById)
	r.GET("/user/:id", controllers.GetReservationsByUserId)
	r.DELETE("/:id", controllers.DeleteReservation)
	r.POST("/approve/:id", controllers.ApproveReservation)
	r.POST("/reject/:id", controllers.RejectReservation)

	r.Run(":8080")
}
