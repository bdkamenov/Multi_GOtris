package main

import (
	"github.com/bdkamenov/tetris_multiplayer/core"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().Unix())

	ebiten.Run(core.Update, 800, 600, 1, "Tetris")
}