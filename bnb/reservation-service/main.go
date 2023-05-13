package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
)

func main() {
	repos.Connect()
	ginSetup()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()

	r.Group("/reservation")

	//TODO: Dodati kontroler endpointove

	r.Run(":8080")
}
