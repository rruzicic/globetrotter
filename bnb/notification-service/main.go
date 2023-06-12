package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	grpc_server "github.com/rruzicic/globetrotter/bnb/accommodation-service/grpc_server"
	"github.com/rruzicic/globetrotter/bnb/notification-service/controllers"
	"github.com/rruzicic/globetrotter/bnb/notification-service/middleware"
	"github.com/rruzicic/globetrotter/bnb/notification-service/repos"
)

//define buffer sizes
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
//define map for assigning connections to users
var connectedClients = struct {
	sync.RWMutex
	clients map[string]*websocket.Conn
}{clients: make(map[string]*websocket.Conn)}

func main() {
	repos.Connect()
	go ginSetup()
	grpc_server.InitServer()
	repos.Disconnect()
}

func enableWebSocketMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		c.Next()
	}
}

func handleWebSocket(c *gin.Context) {
	email := c.Param("email")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	//purely to make sure the connection opened properly
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
	if err != nil {
		log.Println("Error sending message:", err)
	}

	defer conn.Close();

	connectedClients.Lock()
	connectedClients.clients[email] = conn
	connectedClients.Unlock()

	//probably won't be used, left here as an option
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Println("********************")
		log.Printf("Received message from user: %s\n", email)
		log.Printf("Message: %s\n", message)
		log.Println("********************")
	}

	connectedClients.Lock()
	delete(connectedClients.clients, email)
	connectedClients.Unlock()
}

func sendNotification(title string, message string, userEmail string) {
	notificationString := fmt.Sprintf(`{"message": "%s", "title": "%s"}`, message, title)
	notification := []byte(notificationString)
	connectedClients.RLock()
	for key, client := range connectedClients.clients {
		if(key == userEmail) {
			if client.WriteMessage(websocket.TextMessage, []byte(notification)) != nil {
				log.Println("Error sending message to client:")
			}
		}
	}
connectedClients.RUnlock()
}

func ginSetup() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	r.Use(enableWebSocketMiddleware())

	notification := r.Group("/notification")
	
	notification.GET("/:id", controllers.GetReservationById)
	notification.GET("/health", controllers.HealthCheck)
	notification.GET("/websocket", handleWebSocket)
	r.NoRoute()
	r.Run(":8080")
}