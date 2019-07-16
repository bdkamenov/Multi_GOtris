package server

import (
	"encoding/gob"
	"fmt"
	"github.com/bdkamenov/Multi_GOtris/core"
	"math/rand"
	"net"
	"os"
	"time"
)

type player struct {
	Id      int
	Encoder gob.Encoder
	Decoder gob.Decoder
}

var players []player

// StartServer starts the first player and the server side of
// the server-client service
func StartServer(serverIP string, numOfPlayers int) {
	service := serverIP + ":1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	fmt.Println(tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	players = make([]player, 0, numOfPlayers)

	fmt.Println(players)

	println(numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		defer conn.Close()

		fmt.Println("server connection made")
		players = append(players, player{i, *gob.NewEncoder(conn), *gob.NewDecoder(conn)})
	}

	seed := time.Now().Unix()
	rand.Seed(seed)

	for i := 0; i < numOfPlayers; i++ {

		err = players[i].Encoder.Encode(seed)
		checkError(err)
	}

	playersInfo := make([]core.Player, numOfPlayers)

	for i := 0; i < numOfPlayers; i++ {

		err = players[i].Decoder.Decode(&playersInfo[i])
		checkError(err)

		for j := 0; j < numOfPlayers; j++ {
			if j != i {
				err = players[j].Encoder.Encode(playersInfo[i])
				checkError(err)
			}
		}
	}

	go func() {
		for i := 0; i < numOfPlayers; i++ {

			err = players[i].Decoder.Decode(&playersInfo[i])
			checkError(err)

			for j := 0; j < numOfPlayers; j++ {
				if j != i {
					err = players[j].Encoder.Encode(playersInfo[i])
					checkError(err)
				}
			}
		}
		return
		//os.Exit(0)
	}()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
