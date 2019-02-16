package main

import (
	"flag"
	"fmt"
	"github.com/bdkamenov/Multi_GOtris/client"
	"github.com/bdkamenov/Multi_GOtris/core"
	"github.com/bdkamenov/Multi_GOtris/server"
	"math/rand"
	"time"
)

func main() {

	startServer := flag.Bool("server", false, "start game server")
	serverIP := flag.String("connect", "127.0.0.1", "start client connecting to ip")
	playerName := flag.String("name", "player", "enter player name")
	gameMode := flag.String("mode", "", "game mode: time-attack, classic, single")

	flag.Parse()

	if *gameMode == "" {
		println("Please choose game mode!")
		return
	}

	if *gameMode == "single" {
		rand.Seed(time.Now().Unix())
		core.StartGame(*gameMode)
	}

	if *startServer {
		fmt.Println("Server starting")
		server.StartServer(*serverIP, *playerName, *gameMode)
	} else if *gameMode != "single" {
		fmt.Println("Client starting")
		client.StartClient(*serverIP, *playerName, *gameMode)
	}

}
