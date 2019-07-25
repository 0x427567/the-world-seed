package main

import (
	"github.com/0x427567/the-world-seed/pkg/server"
)

func main() {
	server.Config.Port = ":5000"
	server.Run()
}
