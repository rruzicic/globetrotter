package main

import (
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/repos"
)

func main() {
	repos.Connect()
	repos.Disconnect()
}
