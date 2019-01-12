package core

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

var board Board

var score int
var clearedRows int
var level int
var levelUpRate int

func (board *Board) addShape(shape Shape) {
	for i := 0; i < ShapePieces; i++ {
		board[shape.points[i].Y][shape.points[i].X] = shape.color
	}
	score += 10
}

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
		score += 100
	} else if clearedAtOnes < 4 {
		score += 150 * clearedAtOnes // combo cleared lines bonus
	} else if clearedAtOnes == 4 {
		score += 200 * clearedAtOnes // combo cleared lines bonus
	}
	clearedRows+=clearedAtOnes
}