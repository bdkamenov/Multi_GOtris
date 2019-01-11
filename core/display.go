package core

import (
	eb "github.com/hajimehoshi/ebiten"
	ebutil "github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
	"math/rand"
	"time"
)

var screenWidth = 800
var screenHeight = 600

const del = 0.8

var last = time.Now()

var timer = 0.0
var delay = del
var firstRun = true

//var colorNum = rand.Int() % 7
var colorNum = 0


func inBoardCheck() bool {
	for i := 0; i < 4; i++ {
		if ShapeA[i].X < 0 || ShapeA[i].X >= BoardCols || ShapeA[i].Y >= BoardRows {
			return false
		} else if Board[ShapeA[i].Y][ShapeA[i].X] != 0 {
			return false
		}
	}

	return true
}

func Update(screen *eb.Image) error {

	// Perform time processing events
	dt := time.Since(last).Seconds()
	last = time.Now()
	timer += dt

	rotate := false
	dx := 0

	backGr, _, _ := ebutil.NewImageFromFile("assets/image/tetris_backgraund.png", eb.FilterDefault)
	//fmt.Println(err)

	if eb.IsKeyPressed(eb.KeyUp) == true {
		rotate = true
	} else if eb.IsKeyPressed(eb.KeyLeft) {
		dx = -1
	} else if eb.IsKeyPressed(eb.KeyRight) {
		dx = 1
	} else if eb.IsKeyPressed(eb.KeyDown) {
		delay = 0.05
	}

	if firstRun {
		colorNum = rand.Int()%7 + 1

		for i := 0; i < 4; i++ {
			ShapeA[i].X = figures[colorNum][i] % 2
			ShapeA[i].Y = figures[colorNum][i] / 2
		}

		firstRun = false
	}

	/// <- Move -> ///

	copy(BufferB[:], ShapeA[:]) // save the positions before inBoardCheck
	for i := 0; i < 4; i++ {
		ShapeA[i].X += dx
	}
	if !inBoardCheck() { // if the the position is not in the board go back
		copy(ShapeA[:], BufferB[:])
	}

	/// Rotate ///
	if rotate {
		centerOfRot := ShapeA[1] // center of rotation

		for i := 0; i < 4; i++ {
			x := ShapeA[i].Y - centerOfRot.Y
			y := ShapeA[i].X - centerOfRot.X
			ShapeA[i].X = centerOfRot.X - x
			ShapeA[i].Y = centerOfRot.Y + y
		}

		if !inBoardCheck() { // if there is no where to go we dont move the figure again
			copy(ShapeA[:], BufferB[:])
		}

	}

	/// Tick ///

	if timer > delay {
		copy(BufferB[:], ShapeA[:]) // save the positions before inBoardCheck
		for i := 0; i < 4; i++ {
			ShapeA[i].Y += 1
		}

		if !inBoardCheck() {
			for i := 0; i < 4; i++ {
				Board[BufferB[i].Y][BufferB[i].X] = colorNum
			}

			colorNum = rand.Int()%7 + 1
			for i := 0; i < 4; i++ {
				ShapeA[i].X = figures[colorNum][i] % 2
				ShapeA[i].Y = figures[colorNum][i] / 2
			}
		}

		timer = 0
	}

	/// inBoardCheck lines ///

	k := BoardRows - 1
	for i := BoardRows - 1; i > 0; i-- {
		cnt := 0
		for j := 0; j < BoardCols; j++ {
			if Board[i][j] != 0 {
				cnt++
			}

			Board[k][j] = Board[i][j]
		}

		if cnt < BoardCols {
			k--
		}
	}

	rotate = false
	dx = 0
	delay = del

	/// draw ///

	screen.Clear()
	screen.DrawImage(backGr, &eb.DrawImageOptions{})

	for i := 0; i < BoardRows; i++ {
		for j := 0; j < BoardCols; j++ {
			if Board[i][j] == 0 {
				continue
			}
			ebImg, _, _ := ebutil.NewImageFromFile("assets/image/"+colors[Board[i][j] - 1]+".png", eb.FilterDefault)
			geo1 := eb.GeoM{}
			geo1.Translate(float64(j*25+299), float64(i*25+95))
			screen.DrawImage(ebImg, &eb.DrawImageOptions{GeoM: geo1})
		}
	}

    println(colorNum)
	println(colors[colorNum - 1])

	for i := 0; i < 4; i++ {
		ebImg, _, _ := ebutil.NewImageFromFile("assets/image/"+colors[colorNum - 1]+".png", eb.FilterDefault)
		geo1 := eb.GeoM{}
		geo1.Translate(float64(ShapeA[i].X*25+299), float64(ShapeA[i].Y*25+95))
		screen.DrawImage(ebImg, &eb.DrawImageOptions{GeoM: geo1})
	}

	return nil
}
