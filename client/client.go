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

// StartClient starts the game for the second player and
// the client side for the server-client service

func StartClient(serverIP, playerName, mode string, numOfPlayers int) {

	conn, err := net.Dial("tcp", serverIP+":1234")
	checkError(err)

	fmt.Println("client connection made")

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	time.Sleep(3 * time.Second)

	var seed int64
	err = decoder.Decode(&seed)
	checkError(err)
	rand.Seed(seed)

	core.Player1 = core.Player{Name:playerName}

	err = encoder.Encode(core.Player1)
	checkError(err)

	core.OtherPlayers = make([]core.Player, 0, numOfPlayers - 1)

	for _, v := range core.OtherPlayers {
		err = decoder.Decode(&v)
		checkError(err)
	}

	go func() {

		for {
			encoder.Encode(core.Player1)

			for _, v := range core.OtherPlayers {
				err = decoder.Decode(&v)
				checkError(err)
				if v.GameOver == true {
					println(v.Name, "lost, you win!")
					break
				}
			}

			//time.Sleep(100 * time.Millisecond)
		}

		os.Exit(0)
	}()

	core.StartGame(mode)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
