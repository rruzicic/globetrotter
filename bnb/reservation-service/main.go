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

	res := r.Group("/reservation")
	res.POST("/", controllers.CreateReservation)
	res.GET("/:id", controllers.GetReservationById)
	res.GET("/accommodation/:id", controllers.GetReservationsByAccommodationId)
	res.GET("/user/:id", controllers.GetReservationsByUserId)
	res.DELETE("/:id", controllers.DeleteReservation)
	res.POST("/approve/:id", controllers.ApproveReservation)
	res.POST("/reject/:id", controllers.RejectReservation)
	res.POST("/accommodation/:acc_id/reservation/:res_id", controllers.AddReservationToAccommodation)
	res.GET("/test/:msg", controllers.TestConnection)

	r.Run(":8080")
}
