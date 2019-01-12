package client

import (
	"bufio"
	"encoding/gob"
	"fmt"
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
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		newEntity := text
		err := encoder.Encode(newEntity)
		checkError(err)
		var entity string
		err = decoder.Decode(&entity)
		checkError(err)
		fmt.Println("client receive data: ",entity)
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}