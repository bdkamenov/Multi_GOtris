package core

import (
	"math/rand"
)

const ShapePieces = 4

type Piece = int

const (
	IShape Piece = iota + 1
	JShape
	LShape
	OShape
	SShape
	TShape
	ZShape
	ShapesCount
)

// movement directions
const (
	Left int = iota - 1
	Center
	Right  // also used for down
)

type Point struct {
	X, Y int
}

// Shape is a type containing four points, which represents the four points
// making a contiguous 'piece'.
type Shape struct {
	points [ShapePieces]Point
	color  Color
}

var ActiveShape Shape
var NextShape Shape
var HoldedShape Shape
var Buffer Shape

// every figure can be described in
// 2x4 matrix and every matrix elements(figure)
// can be described using the idecies of matrices
var figures = [ShapesCount][ShapePieces]int{
	{0, 0, 0, 0}, // empty
	{1, 3, 5, 7}, // I
	{2, 4, 5, 7}, // Z
	{3, 5, 4, 6}, // S
	{3, 5, 4, 7}, // T
	{2, 3, 5, 7}, // L
	{3, 5, 7, 6}, // J
	{2, 3, 4, 5}} // O

var colors = [7]string{"I", "Z", "S", "T", "L", "J", "O"}

func holdShape() {
	if HoldedShape.color == Empty {
		ActiveShape.resetShape()
		HoldedShape.copyFrom(&ActiveShape)
		ActiveShape.copyFrom(&NextShape)
		ActiveShape.moveShape(4, Center)
		NextShape = generateNewShape()
	} else {
		var temp Shape
		ActiveShape.resetShape()
		temp.copyFrom(&ActiveShape)
		ActiveShape.copyFrom(&HoldedShape)
		HoldedShape.copyFrom(&temp)
		ActiveShape.moveShape(4, Center)
	}
}

func (shape *Shape) rotate() {

	if shape.color == 7 {
		return
	}

	centerOfRot := shape.points[1] // center of rotation

	var buff Shape
	buff.copyFrom(shape)

	for i := 0; i < ShapePieces; i++ {
		x := buff.points[i].Y - centerOfRot.Y
		y := buff.points[i].X - centerOfRot.X
		buff.points[i].X = centerOfRot.X - x
		buff.points[i].Y = centerOfRot.Y + y
	}

	if !buff.isInside(gameBoard) {
		buff.moveShape(1, 0)

		if buff.isInside(gameBoard) {
			shape.copyFrom(&buff)
			return
		}

		buff.moveShape(-2, 0)

		if buff.isInside(gameBoard) {
			shape.copyFrom(&buff)
			return
		}
	} else {
		shape.copyFrom(&buff)
	}
}

func (shape *Shape) moveLeftRight(direction int, buffer *Shape) {

	buffer.copyFrom(shape) // save the positions before inBoardCheck
	shape.moveShape(direction, Center)
}

func (shape *Shape) copyFrom(other *Shape) {
	copy(shape.points[:], other.points[:])
	shape.color = other.color
}

func (shape *Shape) isInside(board Board) bool {
	for i := 0; i < ShapePieces; i++ {
		if shape.points[i].X < 0 || shape.points[i].X >= BoardCols || shape.points[i].Y >= BoardRows {
			return false
		} else if board[shape.points[i].Y][shape.points[i].X] != Empty {
			return false
		}
	}

	return true
}

func (shape *Shape) moveShape(r, c int) {
	for i := 0; i < ShapePieces; i++ {
		shape.points[i].X += r
		shape.points[i].Y += c
	}
}

func (shape *Shape) applyGravity() {
	shape.moveShape(Center, Right)
}

func (shape *Shape) resetShape() {

	for i := 0; i < ShapePieces; i++ {
		shape.points[i].X = figures[shape.color][i] % 2
		shape.points[i].Y = figures[shape.color][i] / 2
	}
}

func generateNewShape() Shape {

	colorNum := rand.Int()%7 + 1

	var newShape Shape
	for i := 0; i < ShapePieces; i++ {
		newShape.points[i].X = figures[colorNum][i] % 2
		newShape.points[i].Y = figures[colorNum][i] / 2
	}

	newShape.color = colorNum
	return newShape
}
