# Multi GOtris
A multiplyer golang tetris.<br>
University project for programming with Go course.

### To get the game
* The only dependency the project has is cross platform [Ebiten](https://github.com/hajimehoshi/ebiten)
2D game library used for the graphics, texts, inputs and sounds.
Installing the library is shown for different platforms.<br>
[Here](https://godoc.org/github.com/hajimehoshi/ebiten) is th documentation for the library.

* To get the game run `go get github.com/bdkamenov/Multi_GOtris` or clone the repository. The `go get` command will resolve the dependencies and will get all needed libs to run the game.

* The game uses command-line arguments for running and connecting to 
server and choose the game mode.

### Example run 

* `go run main.go -server -mode=classic -name=the_killer` to run the server
* `go run main.go -connect=127.0.0.1 -name=the_victim` to run the client
* `go run main.go -help` to see all options with the default values

### Controls 
* Left/Right arrow - Move piece
* Up arrow - Rotate piece
* Down arrow - Fast fall
* Space - Hold piece

### Features
* The connection is using tcp protocol for transferring data.

## TODO:
* make the server multithread so more players can play
* fix drawing bugs 

