package server

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type Player struct {
	Name     string
	Score    int
	GameOver bool
}

var players map[string]Player
var seed int64
var allReady bool
var mutex sync.Mutex
var wg sync.WaitGroup

func handlePlayer(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	fmt.Println("server connection made")

	fmt.Println("Wait for all players to join!")
	for !allReady {
	}

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	err := encoder.Encode(seed)
	checkError(err)
	time.Sleep(3 * time.Second)

	player := Player{"", 0, false}

	err = decoder.Decode(&player)
	checkError(err)

	mutex.Lock()
	fmt.Println("decoded: ", player)
	players[player.Name] = player
	mutex.Unlock()

	for {
		mutex.Lock()
		encoder.Encode(players)
		mutex.Unlock()

		mutex.Lock()
		err = decoder.Decode(&players)
		mutex.Unlock()
		checkError(err)

	}

	os.Exit(0)
}

// StartServer starts the first player and the server side of
// the server-client service
func StartServer(serverIP string, numOfPlayers int) {
	service := serverIP + ":1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	fmt.Println(tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	players = make(map[string]Player)
	allReady = false
	seed = time.Now().Unix()

	for i := 0; i < numOfPlayers; i++ {
		wg.Add(1)

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handlePlayer(conn)

	}
	allReady = true

	wg.Wait()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
