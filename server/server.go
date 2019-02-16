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

// StartServer starts the first player and the server side of
// the server-client service

func StartServer(serverIP, playerName, mode string) {
	service := serverIP+":1234"
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
		rand.Seed(seed)

		err = encoder.Encode(seed)
		checkError(err)

		core.Player1 = core.Player{playerName, 0, false}

		err = decoder.Decode(&core.Player2)
		checkError(err)

		go func() {
			for {
				encoder.Encode(core.Player1)

				err = decoder.Decode(&core.Player2)
				checkError(err)

				if core.Player2.GameOver == true {
					println(core.Player2.Name, "lost, you win!")
					break
				}

			}

			conn.Close()
			os.Exit(0)
		}()

		// start game here
		core.StartGame(mode)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
