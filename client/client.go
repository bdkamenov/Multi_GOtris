package client

import (
	"encoding/gob"
	"fmt"
	"github.com/bdkamenov/Multi_GOtris/core"
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

	time.Sleep(3*time.Second)

	var seed int
	err = decoder.Decode(&seed)
	checkError(err)
	fmt.Println("client receive seed: ", seed)

	player := core.Player{playerName, 0, false}

	for  {

		fmt.Println("client send Player", player.Name, player.Score)
		encoder.Encode(player)



		var otherPlayer core.Player
		err := decoder.Decode(&otherPlayer)
		checkError(err)
		fmt.Println("client recieve Player", otherPlayer.Name, otherPlayer.Score)

		//updateGame
		player.Score++


		//rand.Seed(time.Now().Unix())
		//core.SetupScene()
		//ebiten.Run(core.Update, 800, 600, 1, "Tetris")
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}