# Multi GOtris
A multiplyer golang tetris.<br>
University project for programming with Go course.

### Run the game
* The only dependency the project has is cross platform [Ebiten](https://github.com/hajimehoshi/ebiten)
2D game library used for the graphics, texts, inputs and sounds.
Installing the library is shown for different platforms.<br>
[Here](https://godoc.org/github.com/hajimehoshi/ebiten) is th documentation for the library.

* To get the game run `go get github.com/bdkamenov/Multi_GOtris` or clone the repository.

* The game uses command-line arguments for running and connecting to 
server and choose the game mode.

### Controls 
* Left/Right arrow - Move piece
* Up arrow - Rotate piece
* Down arrow - Fast fall
* Space - Hold piece

### Features
* The connection is using tcp connection.

### Example run 

* `go run main.go`

##TODO:
* Make communication between the players(many clients with the server)
* command line example ./gotris -connect 127.0.0.1 -mode classic
