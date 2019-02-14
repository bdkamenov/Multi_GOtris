package server

import (
	"encoding/gob"
	"fmt"
	"github.com/bdkamenov/Multi_GOtris/core"
	"net"
	"os"
	"time"
)

func StartServer(playerName string) {
	service := "127.0.0.1:1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	fmt.Println(tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("server connection made")

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		seed := time.Now().Unix()
		err = encoder.Encode(seed)
		checkError(err)

		println("Seed sended: ", seed)
		player := core.Player{playerName, 0, false}

		for {
			var otherPlayer core.Player
			err = decoder.Decode(&otherPlayer)
			checkError(err)
			fmt.Println("server recieved other Player data: ", otherPlayer.Name, otherPlayer.Score)

			fmt.Println("server send data", player.Name, player.Score)
			encoder.Encode(player)

			/// update game
			player.Score++
			time.Sleep(3*time.Second)
		}

		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}