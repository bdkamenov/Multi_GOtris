
package main

import (
	"fmt"
	"time"

	"github.com/bdkamenov/tetris_multiplayer/client"
	"github.com/bdkamenov/tetris_multiplayer/server"
)

func main() {
	go func() {
		fmt.Println("Server starting")
		server.StartServer()
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Client starting")
	client.StartClient()
}