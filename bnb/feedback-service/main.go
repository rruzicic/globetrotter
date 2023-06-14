package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/middlewares"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"
)

func main() {
	repos.Connect()
	go ginSetup()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())
	r.NoRoute()

	//TODO implement controllers

	r.Run(":8080")
}
