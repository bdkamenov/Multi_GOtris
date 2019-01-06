package server

import (
	"encoding/gob"
	"fmt"
	"github.com/bdkamenov/tetris_multiplayer/core"
	"net"
	"os"
)

func StartServer() {
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

		for n := 0; n < 10; n++ {
			var entity core.Point
			err = decoder.Decode(&entity)
			checkError(err)
			fmt.Println("server receive data: ", entity.X, " ", entity.Y)
			fmt.Println("server send data")
			entity = core.Point{float32(n*1), float32(n*2)}
			err = encoder.Encode(entity)
			checkError(err)
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