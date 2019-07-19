package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/bdkamenov/Multi_GOtris/client"
	"github.com/bdkamenov/Multi_GOtris/core"
	"github.com/bdkamenov/Multi_GOtris/server"
)

func main() {

	startServer := flag.Bool("server", false, "start game server")
	serverIP := flag.String("connect", "127.0.0.1", "start client connecting to ip")
	playerName := flag.String("name", "player", "enter player name")
	gameMode := flag.String("mode", "", "game mode: time-attack, classic, single")
	numOfPlayers := flag.Int("numofplayers", 1, "number of the players playing together")

	flag.Parse()

	if !(*startServer) && *gameMode == "" {
		println("Please choose game mode!")
		return
	}

	if *gameMode == "single" {
		rand.Seed(time.Now().Unix())
		core.StartGame(*gameMode)
	}

	if *startServer {
		fmt.Println("Server starting")
		server.StartServer(*serverIP, *numOfPlayers)
	} else if *gameMode != "single" {
		fmt.Println("Client starting")
		client.StartClient(*serverIP, *playerName, *gameMode, *numOfPlayers)
	}

}
