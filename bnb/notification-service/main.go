package main

import (
	"github.com/gin-gonic/gin"
	grpc_server "github.com/rruzicic/globetrotter/bnb/accommodation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/notification-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/notification-service/middleware"
	"github.com/rruzicic/globetrotter/bnb/notification-service/repos"
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
	r.Use(middleware.CORSMiddleware())

	notification := r.Group("/notification")
	
	notification.GET("/:id", controllers.GetReservationById)
	notification.GET("/health", controllers.HealthCheck)
	r.NoRoute()
	r.Run(":8080")
}