package server

import (
	"github.com/gin-gonic/gin"
	"github.com/0x427567/the-world-seed/pkg/websocket"
)

func Run() {
	router := gin.Default()
	router.GET("/websocket", websocket.Handle)
	router.Run(":5000")
}
