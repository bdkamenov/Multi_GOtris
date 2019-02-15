package core

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	eb "github.com/hajimehoshi/ebiten"
	ebutil "github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	_ "image/png"

	"io/ioutil"

	"math"
	"time"
)

var delayBuffer float64

var last time.Time
var timer float64
var delay = delayBuffer

var textures [8]*eb.Image
var linesScoreFont font.Face
var holdNextFont font.Face

var leftRightDelay float64
var moveCounter int
var rotateHoldDelay bool

func loadTextures() {
	for i := 0; i < 7; i++ {

		textures[i], _, _ = ebutil.NewImageFromFile("assets/image/"+colors[i]+".png", eb.FilterDefault)
	}
	textures[7], _, _ = ebutil.NewImageFromFile("assets/image/tetris_backgraund.png", eb.FilterDefault)
}

func StartGame(mode string) {

	SetupScene(mode)
	err := eb.Run(Update, 800, 600, 1, "Tetris")

	if err != nil {
		println("Fatal error: ", err)
	}

}

func SetupScene(mode string) {

	delayBuffer = 0.8
	timer = 0.0

	if mode == timeAttack {
		start = time.Now()
	}

	loadTextures()

	dat, _ := ioutil.ReadFile("assets/fonts/tetris.ttf")
	tt, _ := truetype.Parse(dat)

	linesScoreFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	holdNextFont = truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	last = time.Now()
	ActiveShape = generateNewShape()
	ActiveShape.moveShape(4, Center)
	NextShape = generateNewShape()
	levelUpRate = 300
	level = 0
}

func DrawPiece(piece Point, screen *eb.Image, color int, trX, trY, scX, scY float64) {
	geo1 := eb.GeoM{}
	geo1.Translate(float64(piece.X*25)+trX, float64(piece.Y*25)+trY)
	geo1.Scale(scX, scY)
	screen.DrawImage(textures[color], &eb.DrawImageOptions{GeoM: geo1})
}

func DrawShape(shape Shape, screen *eb.Image, trX, trY, scX, scY float64) {

	for i := 0; i < ShapePieces; i++ {
		DrawPiece(shape.points[i], screen, shape.color-1, trX, trY, scX, scY)
	}
}

func DrawText(score, lines int, screen *eb.Image) {

	scoreText := fmt.Sprintf("Score: %d", score)
	linesText := fmt.Sprintf("Lines: %d", lines)
	levelText := fmt.Sprintf("Level: %d", level)
	text.Draw(screen, scoreText, linesScoreFont, 295, 78, color.White)
	text.Draw(screen, linesText, linesScoreFont, 70, 100, color.White)
	text.Draw(screen, "Next Piece", holdNextFont, 605, 74, color.White)
	text.Draw(screen, "Holded", holdNextFont, 605, 277, color.White)
	text.Draw(screen, levelText, linesScoreFont, 595, 425, color.White)

	text.Draw(screen, "Players:", linesScoreFont, 70, 200, color.White)
	playerText := fmt.Sprintf("%d", Player2.Score)

	text.Draw(screen, Player2.Name+": "+playerText, linesScoreFont, 70, 260, color.White)

}

func DrawBoard(board Board, screen *eb.Image) {

	for i := 0; i < BoardRows; i++ {
		for j := 0; j < BoardCols; j++ {
			if board[i][j] == 0 {
				continue
			}
			DrawPiece(Point{j, i}, screen, board[i][j]-1, 299, 95, 1, 1)
		}
	}
}

func GetInput() (rotate bool, direction int) {
	if eb.IsKeyPressed(eb.KeyUp) && rotateHoldDelay == false {
		rotateHoldDelay = true
		rotate = true
	} else if eb.IsKeyPressed(eb.KeyLeft) && leftRightDelay == 0.0 {
		if moveCounter > 0 {
			leftRightDelay = 0.05
		} else {
			leftRightDelay = 0.3
		}
		moveCounter++
		direction = Left
	} else if eb.IsKeyPressed(eb.KeyRight) && leftRightDelay == 0.0 {
		if moveCounter > 0 {
			leftRightDelay = 0.05
		} else {
			leftRightDelay = 0.3
		}
		moveCounter++
		direction = Right
	} else if eb.IsKeyPressed(eb.KeyDown) {
		delayBuffer = delay
		delay = 0.05
	} else if eb.IsKeyPressed(eb.KeySpace) && rotateHoldDelay == false {
		holdShape()
		rotateHoldDelay = true
	}

	return
}

func Update(screen *eb.Image) error {

	// Perform time processing events
	dt := time.Since(last).Seconds()
	last = time.Now()
	timer += dt

	rotate := false
	direction := Center

	if leftRightDelay > 0.0 {
		leftRightDelay = math.Max(leftRightDelay-dt, 0.0)
	}

	rotate, direction = GetInput()

	if !eb.IsKeyPressed(eb.KeySpace) && !eb.IsKeyPressed(eb.KeyUp) {
		rotateHoldDelay = false
	}

	if !eb.IsKeyPressed(eb.KeyRight) && !eb.IsKeyPressed(eb.KeyLeft) {
		moveCounter = 0.0
		leftRightDelay = 0.0
	}

	/// <- Move -> ///
	ActiveShape.moveLeftRight(direction, &Buffer)

	if !ActiveShape.isInside(gameBoard) { // if the the position is not in the gameBoard go back
		ActiveShape.copyFrom(&Buffer)
	}

	/// Rotate ///
	if rotate {
		ActiveShape.rotate()

		if !ActiveShape.isInside(gameBoard) { // if there is no where to go we dont moveLeftRight the figure again
			ActiveShape.copyFrom(&Buffer)
		}
	}

	/// Tick ///
	if timer > delay {
		Buffer.copyFrom(&ActiveShape) // save the positions before inBoardCheck
		ActiveShape.applyGravity()

		if !ActiveShape.isInside(gameBoard) {
			gameBoard.addShape(Buffer)

			if gameBoard.isGameOver() {
				Player1.GameOver = true
				println("game over")
				return nil
			}

			ActiveShape = NextShape
			ActiveShape.moveShape(4, Center)
			NextShape = generateNewShape()
		}

		timer = 0
	}

	/// clear lines ///
	gameBoard.clearLines()

	rotate = false

	/// level up if needed ///
	if levelUpRate < Player1.Score && delay > 0.2 {
		delay -= 0.1
		delayBuffer = delay
		levelUpRate += Player1.Score
		level++
	}
	delay = delayBuffer

	/// draw ///

	screen.Clear()
	screen.DrawImage(textures[7], &eb.DrawImageOptions{})
	DrawBoard(gameBoard, screen)
	DrawShape(ActiveShape, screen, 299, 95, 1, 1)
	DrawShape(NextShape, screen, 612, 90, 1.05, 1.05)
	DrawText(Player1.Score, clearedRows, screen)
	if HoldedShape.color != Empty {
		DrawShape(HoldedShape, screen, 780, 350, 0.8, 0.8)
	}

	return nil
}
