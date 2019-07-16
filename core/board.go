package core

import (
	"encoding/gob"
	"time"
)

const timeAttack = "time-attack"


// Player describes the player nickname, score if its
type Player struct {
	Name     string
	Score    int
	GameOver bool
	Encoder gob.Encoder
	Decoder gob.Decoder
}

var start time.Time

type Color = int // represents the color of the piece

const (
	Empty Color = iota
	Blue
	Purple
	Green
	Red
	Orange
	Yellow
	SkyBlue
)

const BoardRows = 17
const BoardCols = 10

type Board [BoardRows][BoardCols]Color

var gameBoard Board

var clearedRows int
var level int
var levelUpRate int

var Player1 Player
var OtherPlayers []Player

// addShape inserts shape in the board
func (board *Board) addShape(shape Shape) {

	for i := 0; i < ShapePieces; i++ {
		if shape.points[i].X < BoardCols && shape.points[i].Y < BoardRows {
			board[shape.points[i].Y][shape.points[i].X] = shape.color
		}
	}
	Player1.Score += 10
}

// clearLines checks if there are fitted lines and removes them
func (board *Board) clearLines() {
	clearedAtOnes := 0
	k := BoardRows - 1
	for i := BoardRows - 1; i > 0; i-- {
		cnt := 0
		for j := 0; j < BoardCols; j++ {
			if board[i][j] != 0 {
				cnt++
			}
			board[k][j] = board[i][j]
		}

		if cnt < BoardCols {
			k--
		} else {
			clearedAtOnes++
		}
	}

	if clearedAtOnes == 1 {
		Player1.Score += 100
	} else if clearedAtOnes < 4 {
		Player1.Score += 150 * clearedAtOnes // combo cleared lines bonus
	} else if clearedAtOnes == 4 {
		Player1.Score += 200 * clearedAtOnes // combo cleared lines bonus
	}
	clearedRows += clearedAtOnes
}

// isGameOver checks if there is a Piece in the last line of the board
// and ends the game if there is
func (board Board) isGameOver() bool {

	var zero time.Time

	if start != zero && time.Now().Sub(start) > time.Minute {
		return true
	}

	for i := 0; i < BoardCols; i++ {
		if board[1][i] != 0 {
			return true
		}
	}
	return false
}
