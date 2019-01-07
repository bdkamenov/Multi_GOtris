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

const M = 16
const N = 10

var Board [M][N]int

type Point struct {
	X, Y int
}

var a, b [4]Point

// every figure can be described in
// 2x4 matrix and every matrix elements(figure)
// can be described using the idecies of matrices
var figures = [7][4]int{
	{1, 3, 5, 7}, // I
	{2, 4, 5, 7}, // Z
	{3, 5, 4, 6}, // S
	{3, 5, 4, 7}, // T
	{2, 3, 5, 7}, // L
	{3, 5, 7, 6}, // J
	{2, 3, 4, 5}} // O

var colors = [7]string{"I", "J", "L", "S", "Z", "O", "T"}

///
/// TODO: game timing. Ticker delay and timer !!
///

var last = time.Now()

var timer = 0.0
var delay = 0.7

func check() bool {
	for i := 0; i < 4; i++ {
		if a[i].X < 0 || a[i].X >= N || a[i].Y >= M {
			return false
		} else if Board[a[i].Y][a[i].X] != 0 {
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
	colorNum := 3

	ebImg, _, _ := ebutil.NewImageFromFile("assets/image/"+colors[colorNum]+".png", eb.FilterDefault)
	backGr, _, _ := ebutil.NewImageFromFile("assets/image/tetris_backgraund.png", eb.FilterDefault)
	//fmt.Println(err)

	if eb.IsKeyPressed(eb.KeyUp) == true {
		rotate = true
	} else if eb.IsKeyPressed(eb.KeyLeft) {
		dx = -1
	} else if eb.IsKeyPressed(eb.KeyRight) {
		dx = 1
	}

	/// <- Move -> ///

	for i := 0; i < 4; i++ {
		b[i] = a[i]
		a[i].X += dx
	}
	if !check() {
		for i := 0; i < 4; i++ {
			a[i] = b[i]
		}
	}

	/// Rotate ///

	//fmt.Println(rotate)
	if rotate {
		centerOfRot := a[1] // center of rotation

		for i := 0; i < 4; i++ {
			x := a[i].Y - centerOfRot.Y
			y := a[i].X - centerOfRot.X
			a[i].X = centerOfRot.X - x
			a[i].Y = centerOfRot.Y + y
		}

		if !check() {
			for i := 0; i < 4; i++ {
				a[i] = b[i]
			}
		}

	}

	/// Tick ///

	if timer > delay {
		for i := 0; i < 4; i++ {
			b[i] = a[i]
			a[i].Y += 1
		}

		if !check() {
			for i := 0; i < 4; i++ {
				Board[b[i].Y][b[i].X] = colorNum

				colorNum = 1 + rand.Int()%7
				n := rand.Int() % 7
				for i := 0; i < 4; i++ {
					a[i].X = figures[n][i] % 2
					a[i].Y = figures[n][i] / 2
				}
			}

		}

		timer = 0
	}


	/// draw ///

	screen.Clear()
	screen.DrawImage(backGr, &eb.DrawImageOptions{})


	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if Board[i][j] == 0 {
				continue
			}
			geo1 := eb.GeoM{}
			geo1.Translate(float64(j*25+300), float64(i*25+120))
			screen.DrawImage(ebImg, &eb.DrawImageOptions{GeoM:geo1})
		}
	}


	for i := 0; i < 4; i++ {
		geo1 := eb.GeoM{}
		//geo.Scale(0.2, 0.2)
		geo1.Translate(float64(a[i].X*25+300), float64(a[i].Y*25+120))
		screen.DrawImage(ebImg, &eb.DrawImageOptions{GeoM: geo1})
	}

	rotate = false
	dx = 0
	//ebutil.DebugPrint(screen, "Our first game in Ebiten!")

	return nil
}
