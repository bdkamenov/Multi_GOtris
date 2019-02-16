package core

import "testing"

func TestMoveShape(t *testing.T) {

	shape := generateNewShape()

	var expected Shape
	expected.copyFrom(&shape)

	for i := 0; i<4 ; i++ {
		expected.points[i].X += 1;
		expected.points[i].Y += 1;
	}

	shape.moveShape(1, 1)

	for i := 0; i < 4; i++ {
		if expected.points[i].X != shape.points[i].X ||
			expected.points[i].Y != shape.points[i].Y {
			t.Error("Expected shape is not like moved")
		}
	}
}