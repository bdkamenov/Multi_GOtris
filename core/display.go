package core

import (
	"fmt"
	eb "github.com/hajimehoshi/ebiten"
	ebutil "github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
)

var screenWidth = 800
var screenHeight = 600

const M = 20
const N = 10

type Board [M][N]int

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

func Update(screen *eb.Image) error {

	rotate := false
	dx := 0
	colorNum := 3
	//timer := 0.0
	//delay := 0.3

	ebImg, _, _ := ebutil.NewImageFromFile("assets/image/"+colors[colorNum]+".png", eb.FilterDefault)

	if eb.IsKeyPressed(eb.KeyUp) == true {
		rotate = true
	} else if eb.IsKeyPressed(eb.KeyLeft) {
		dx = -1
	} else if eb.IsKeyPressed(eb.KeyRight) {
		dx = 1
	}

	/// <- Move -> ///

	for i := 0; i < 4; i++ {
		a[i].X += dx
	}

	/// Rotate ///

fmt.Println(rotate)
	if rotate {
		centerOfRot := a[1] // center of rotation

		for i := 0; i < 4; i++ {
			x := a[i].Y - centerOfRot.Y
			y := a[i].X - centerOfRot.X
			a[i].X = centerOfRot.X - x
			a[i].Y = centerOfRot.Y + y
		}

	}

//	time.Sleep(time.Second / 5)

	n := 3
	if a[0].X == 0 {

	for i := 0; i < 4; i++ {
		a[i].X = figures[n][i] % 2
		a[i].Y = figures[n][i] / 2
	}
	}

	//for i := 0; i < 4; i++ {
	//	fmt.Println(a[i].X, " ", a[i].Y)
	//}
	screen.Clear()

	imageOptions := eb.DrawImageOptions{}
	for i := 0; i < 4; i++ {
		geo := eb.GeoM{}
		//geo.Scale(0.2, 0.2)
		geo.Translate(float64(a[i].X*32), float64(a[i].Y*32))
		imageOptions = eb.DrawImageOptions{GeoM: geo}
		screen.DrawImage(ebImg, &imageOptions)
	}

	rotate = false
	dx = 0
	//ebutil.DebugPrint(screen, "Our first game in Ebiten!")

	return nil
}
