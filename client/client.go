package client

import (
	"encoding/gob"
	"fmt"
	"github.com/bdkamenov/Multi_GOtris/core"
	"math/rand"
	"net"
	"os"
	"time"
)

func StartClient(serverIP, playerName string) {

	conn, err := net.Dial("tcp", serverIP+":1234")
	checkError(err)

	fmt.Println("client connection made")

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	time.Sleep(3 * time.Second)

	var seed int64
	err = decoder.Decode(&seed)
	checkError(err)
	fmt.Println("client receive seed: ", seed)
	rand.Seed(seed)

	core.Player1 = core.Player{playerName, 0, false}

	fmt.Println("client send Player", core.Player1.Name, core.Player1.Score)
	encoder.Encode(core.Player1)

	err = decoder.Decode(&core.Player2)
	checkError(err)
	fmt.Println("client recieve Player", core.Player2.Name, core.Player2.Score)


	go func() {

		for {

			fmt.Println("client send Player", core.Player1.Name, core.Player1.Score)
			encoder.Encode(core.Player1)

			err := decoder.Decode(&core.Player2)
			checkError(err)
			fmt.Println("client recieve Player", core.Player2.Name, core.Player2.Score)

			//updateGame
			//time.Sleep(1 * time.Second)
		}
	}()
	core.StartGame()

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
