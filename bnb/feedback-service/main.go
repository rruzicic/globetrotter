package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/controllers"
	grpc_server "github.com/rruzicic/globetrotter/bnb/feedback-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"
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

	//TODO implement controllers
	HostFeedback := r.Group("/HostFeedback")
	HostFeedback.POST("/", controllers.CreateHostReview)
	HostFeedback.GET("/:id", controllers.GetHostReviewById)
	HostFeedback.GET("/:user_id", controllers.GetHostReviewsByUserId)
	HostFeedback.GET("/:host_id", controllers.GetHostReviewsByHostId)
	HostFeedback.DELETE("/:id", controllers.DeleteHostReview)
	HostFeedback.PUT("/", controllers.UpdateHostReview)

	AccommodationFeedback := r.Group("AccommodationFeedback")
	AccommodationFeedback.POST("/", controllers.CreateAccommodationReview)
	AccommodationFeedback.GET("/:id", controllers.GetAccommodationReviewById)
	AccommodationFeedback.GET("/:user_id", controllers.GetAccommodationReviewsByUserId)
	AccommodationFeedback.GET("/:accommodation_id", controllers.GetAccommodationReviewsByAccommodationId)
	AccommodationFeedback.DELETE("/:id", controllers.DeleteAccommodationReview)
	AccommodationFeedback.PUT("/", controllers.UpdateAccommodationReview)

	r.Run(":8080")
}
