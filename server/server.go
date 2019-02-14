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
		rand.Seed(seed)

		err = encoder.Encode(seed)
		checkError(err)

		println("Seed sended: ", seed)
		core.Player1 = core.Player{playerName, 0, false}

		err = decoder.Decode(&core.Player2)
		checkError(err)

		go func() {
			for {

				fmt.Println("server send data", core.Player1.Name, core.Player1.Score)
				encoder.Encode(core.Player1)

				err = decoder.Decode(&core.Player2)
				checkError(err)
				fmt.Println("server recieved other Player data: ", core.Player2.Name, core.Player2.Score)

				/// update game
				time.Sleep(3 * time.Second)
			}

			conn.Close()
		}()


		// start game here
		core.StartGame()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
