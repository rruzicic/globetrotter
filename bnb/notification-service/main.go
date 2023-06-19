package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/rruzicic/globetrotter/bnb/notification-service/controllers"
	grpc_server "github.com/rruzicic/globetrotter/bnb/notification-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/notification-service/middleware"
	"github.com/rruzicic/globetrotter/bnb/notification-service/pb"
	"github.com/rruzicic/globetrotter/bnb/notification-service/repos"
	"github.com/rruzicic/globetrotter/bnb/notification-service/socket"
)

func main() {
	repos.Connect()

	//Subscribe to messages
	log.Println("DOBAR DAN JA BIH DA UDJEM U SUBSCRIBE MESSAGE DEO AKO MOZE")
	conn := Conn()
	defer conn.Close()
	_, err := conn.Subscribe("saga-cancel-reservation-3", func(message *nats.Msg) {
		event := pb.ReservationEvent{}
		err := proto.Unmarshal(message.Data, &event)
		if err == nil {
			//Handle the message
			log.Println("Recieved an event about a canceled reservation")
			err = repos.ReservationCanceled(&event)
			if err == nil {
				//Return sucess messsage
				log.Println("Notification created returning success message")
				err = conn.Publish("saga-cancel-reservation-4", []byte("OK"))
				if err != nil {
					log.Panic(err)
				}
			} else {
				//Return error message
				log.Println("Notification not created returning error message")
				err = conn.Publish("saga-cancel-reservation-4", []byte("ERROR"))
				if err != nil {
					log.Panic(err)
				}
			}

		}
	})
	if err != nil {
		log.Panic(err)
	}

	go ginSetup()
	grpc_server.InitServer()
	repos.Disconnect()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(socket.EnableWebSocketMiddleware())

	notification := r.Group("/notification")

	notification.GET("/health", controllers.HealthCheck)
	notification.GET("/websocket/:id", socket.HandleWebSocket)
	notification.GET("/user/:id", controllers.GetNotificationsByUserId)
	r.NoRoute()
	r.Run(":8080")
}

func Conn() *nats.Conn {
	conn, err := nats.Connect("nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
