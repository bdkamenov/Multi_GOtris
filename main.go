package main

import (
	"flag"
	"fmt"
	"github.com/bdkamenov/Multi_GOtris/client"
	"github.com/bdkamenov/Multi_GOtris/server"
)

func main() {

	singlePlay := flag.Bool("single", false, "starts only the game")
	startServer := flag.Bool("server", false, "start game server")
	serverIP := flag.String("connect", "127.0.0.1", "start client connecting to ip")
	playerName := flag.String("name", "player", "enter player name")
	//gameMode := flag.Bool("time-attack", false, "time-attack game mode")

	flag.Parse()

	//if *singlePlay {
	//
	//	rand.Seed(time.Now().Unix())
	//
	//}

	if *startServer {
		fmt.Println("Server starting")
		server.StartServer(*playerName)
	} else if !*singlePlay {
		fmt.Println("Client starting")
		client.StartClient(*serverIP, *playerName)
	}


}