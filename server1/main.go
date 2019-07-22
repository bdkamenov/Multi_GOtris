package main

import (
	"flag"
	"fmt"
	"server/server"
)

func main() {

	startServer := flag.Bool("server", false, "start game server")
	serverIP := flag.String("connect", "127.0.0.1", "start client connecting to ip")
	numOfPlayers := flag.Int("numofplayers", 1, "number of the players playing together")

	flag.Parse()

	if *startServer {
		fmt.Println("Server starting")
		server.StartServer(*serverIP, *numOfPlayers)
	}

}
