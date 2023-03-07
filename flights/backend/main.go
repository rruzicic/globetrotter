package main

import (
	"github.com/rruzicic/globetrotter/flights/backend/repos"
	"github.com/rruzicic/globetrotter/flights/backend/routers"
)

func main() {

	repos.Setup()
	routers.InitRouter()
	repos.Disconnect()
}
