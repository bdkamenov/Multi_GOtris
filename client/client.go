package client

import (
	"encoding/gob"
	"fmt"
	"github.com/bdkamenov/tetris_multiplayer/core"
	"net"
	"os"
)

func StartClient() {


	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	checkError(err)

	fmt.Println("client connection made")

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		fmt.Println("client send data")
		newEntity := core.Point{X: float32(n+1), Y: float32(n+2)}
		err := encoder.Encode(newEntity)
		checkError(err)
		var entity core.Point
		err = decoder.Decode(&entity)
		checkError(err)
		fmt.Println("client receive data: ", entity.X, " ", entity.Y)
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}