package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/middlewares"
)

func main() {
	go ginSetup()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	rec := r.Group("/recommendation")
	rec.GET("/flights")

	r.Run(":8080")
}
