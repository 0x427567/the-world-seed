package server

import (
	"github.com/0x427567/the-world-seed/pkg/websocket"
	"github.com/gin-gonic/gin"
)

type config struct {
	Port string
}

var Config = config{
	":4444",
}

func Run() {
	router := gin.Default()

	router.GET("/websocket", websocket.Handle)

	router.Run(Config.Port)
}
