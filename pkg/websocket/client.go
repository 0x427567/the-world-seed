package websocket

import (
	"github.com/gorilla/websocket"
)

type Pool {

}

type client struct {
	Conn *websocket.Conn
	send chan []byte
}

type Pool struct {
	clients map[*client]bool
}