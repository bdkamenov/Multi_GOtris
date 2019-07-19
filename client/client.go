package client

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"github.com/bdkamenov/Multi_GOtris/core"
)

var mutex sync.Mutex

// StartClient starts the game for the second player and
// the client side for the server-client service
func StartClient(serverIP, playerName, mode string, numOfPlayers int) {

	conn, err := net.Dial("tcp", serverIP+":1234")
	checkError(err)

	fmt.Println("client connection made")

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	time.Sleep(2 * time.Second)

	var seed int64
	err = decoder.Decode(&seed)
	checkError(err)
	rand.Seed(seed)

	core.Player1 = core.Player{Name: playerName, Score: 0, GameOver: false}
	err = encoder.Encode(core.Player1)
	checkError(err)

	core.OtherPlayers = make(map[string]core.Player)
	time.Sleep(5 * time.Second)

	go func() {

		for {
			mutex.Lock()
			decoder.Decode(&core.OtherPlayers)

			core.OtherPlayers[core.Player1.Name] = core.Player1
			encoder.Encode(core.OtherPlayers)

			for _, v := range core.OtherPlayers {
				if v.GameOver == true {
					println(v.Name, "lost, you win!")
					break
				}
			}
			mutex.Unlock()
			time.Sleep(1 * time.Second)
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
