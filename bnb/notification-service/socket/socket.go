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
	id := c.Param("id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	//TODO: remove later: purely to make sure the connection opened properly
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
	if err != nil {
		log.Println("Error sending message:", err)
	}

	defer conn.Close();

	connectedClients.Lock()
	connectedClients.clients[id] = conn
	connectedClients.Unlock()

	<-c.Request.Context().Done()

	connectedClients.Lock()
	delete(connectedClients.clients, id)
	connectedClients.Unlock()
}

func SendNotification(title string, message string, userId string) {
	notificationString := fmt.Sprintf(`{"message": "%s", "title": "%s"}`, message, title)
	log.Println(notificationString)
	notification := []byte(notificationString)
	connectedClients.RLock()
	//TODO: currently sending to all users
	for _, client := range connectedClients.clients {
			if client.WriteMessage(websocket.TextMessage, []byte(notification)) != nil {
				log.Println("Error sending message to client:")
			}
	}
	connectedClients.RUnlock()
}

