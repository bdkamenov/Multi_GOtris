package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/bdkamenov/tetris_multiplayer/core"
)

func main() {
	//go func() {
	//	fmt.Println("Server starting")
	//	server.StartServer()
	//}()
	//
	//time.Sleep(100 * time.Millisecond)
	//
	//fmt.Println("Client starting")
	//client.StartClient()

	ebiten.Run(core.Update, 800, 600, 1, "Tetris")
}