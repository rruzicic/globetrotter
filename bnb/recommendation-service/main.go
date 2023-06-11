package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/repos"
)

func main() {
	ginSetup()
	repos.Connect()
	repos.Disconnect()
}

func ginSetup() {
	log.Println("Setting up gin")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	rec := r.Group("/recommendation", controllers.SearchFlights)
	rec.GET("/flights")

	r.Run(":8080")
}
