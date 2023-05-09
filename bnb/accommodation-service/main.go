package main

import (
	"log"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/pb"
)

func main() {
	c, ctx := makeUserServiceConnection()
	r, err := c.GetUserById(ctx, &pb.UserRequestId{Id: "asdas"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetUser().FirstName)
}
