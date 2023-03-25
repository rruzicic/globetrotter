package main

import (
	"github.com/rruzicic/globetrotter/flights/backend/config"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
	"github.com/rruzicic/globetrotter/flights/backend/routers"
)

func main() {
	config.InitConfig()
	repos.Setup()
	routers.InitRouter()
	repos.Disconnect()
}
