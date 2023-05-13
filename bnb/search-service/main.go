package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/search-service/controllers"
)

func main() {
	ginSetup();
}

func ginSetup() {
	r := gin.New();
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()

	r.GET("/health", controllers.HealthCheck)
}