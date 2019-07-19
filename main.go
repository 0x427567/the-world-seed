package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin:      func(r *http.Request) bool { return true },
		HandshakeTimeout: time.Duration(time.Second * 5),
	}
)

func handleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("cant upgrade connection:", err)
		return
	}

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			log.Println("cant read message:", err)

			switch err.(type) {
			case *websocket.CloseError:
				return
			default:
				return
			}
		}

		// Skip binary messages
		if msgType != websocket.TextMessage {
			continue
		}
		if err = conn.WriteMessage(msgType, msgData); err != nil {
			return
		}

		log.Printf("incoming message: %s\n", msgData)
	}
}

func main() {
	router := gin.Default()
	router.GET("/websocket", handleWebsocket)
	router.Run(":5000")
}
