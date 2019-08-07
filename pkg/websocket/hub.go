package websocket

type hub struct {
	connections map[*client]bool
	broadcast   chan []byte
	register    chan *client
	unregister  chan *client
}

var h = hub{
	broadcast:   make(chan []byte),
	register:    make(chan *client),
	unregister:  make(chan *client),
	connections: make(map[*client]bool),
}

func (h *hub) Broadcast(s string) {
	if len(s) > 0 {
		h.broadcast <- []byte(s)
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
					// go c.ws.Close()
				}
			}
		}
	}
}
