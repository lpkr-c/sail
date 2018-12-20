package shapes

func abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}
	return a
}

type Point struct {
	X float64
	Y float64
}

type Line struct {
	A Point
	B Point
}

type Rectangle struct {
	A Point
	B Point
}

func (rect Rectangle) Range() (float64, float64) {
	return abs(rect.A.X - rect.B.X), abs(rect.A.Y - rect.B.Y)
}

type Triangle struct {
	A Point
	B Point
	C Point
}
