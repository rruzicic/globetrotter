package socket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func EnableWebSocketMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		c.Next()
	}
}

func HandleWebSocket(c *gin.Context) {
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

func SendNotification(title string, message string, userEmail string) {
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

