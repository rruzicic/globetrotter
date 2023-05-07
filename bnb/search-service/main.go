package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginSetup()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute()
	r.Run(":8080")
}
