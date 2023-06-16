package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/notification-service/controllers"
	grpc_server "github.com/rruzicic/globetrotter/bnb/notification-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/notification-service/middleware"
	"github.com/rruzicic/globetrotter/bnb/notification-service/repos"
	"github.com/rruzicic/globetrotter/bnb/notification-service/socket"
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
	r.Use(socket.EnableWebSocketMiddleware())

	notification := r.Group("/notification")
	
	notification.GET("/health", controllers.HealthCheck)
	notification.GET("/websocket/:id", socket.HandleWebSocket)
	notification.GET("/user/:id", controllers.GetNotificationsByUserId)
	r.NoRoute()
	r.Run(":8080")
}