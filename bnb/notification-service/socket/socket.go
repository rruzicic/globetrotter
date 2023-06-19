package socket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rruzicic/globetrotter/bnb/notification-service/model"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
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

	defer conn.Close();

	connectedClients.Lock()
	connectedClients.clients[id] = conn
	connectedClients.Unlock()

	<-c.Request.Context().Done()

	connectedClients.Lock()
	delete(connectedClients.clients, id)
	connectedClients.Unlock()
}

func SendNotification(notification model.Notification) {
	connectedClients.RLock()

	for clientId, client := range connectedClients.clients {
			if(clientId == notification.UserId) {
				if client.WriteMessage(websocket.TextMessage, []byte(notification.Type)) != nil {
					log.Println("Error sending message to client:")
				}
			}
	}
	connectedClients.RUnlock()
}

