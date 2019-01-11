package core


type Point struct {
	X, Y int
}

const (
	IPiece Piece = iota + 1
	JPiece
	LPiece
	OPiece
	SPiece
	TPiece
	ZPiece
	PieceCount
)

type Piece = int
type Block = int

// Shape is a type containing four points, which represents the four points
// making a contiguous 'piece'.
type Shape [4]Point

var ShapeA, BufferB Shape

// every figure can be described in
// 2x4 matrix and every matrix elements(figure)
// can be described using the idecies of matrices
var figures = [PieceCount][4]int{
	{0, 0, 0, 0}, // empty
	{1, 3, 5, 7}, // I
	{2, 4, 5, 7}, // Z
	{3, 5, 4, 6}, // S
	{3, 5, 4, 7}, // T
	{2, 3, 5, 7}, // L
	{3, 5, 7, 6}, // J
	{2, 3, 4, 5}} // O

var colors = [7]string{"I", "Z", "S", "T", "L", "J", "O"}
