package shapes

import "math"

func abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}
	return a
}

func pointRange(a Point, b Point) (float64, float64) {
	return abs(a.X - b.X), abs(a.Y - b.Y)
}

type Point struct {
	X float64
	Y float64
}

type Line struct {
	A Point
	B Point
}

func (l Line) Distance() float64 {
	xDiff, yDiff := pointRange(l.A, l.B)
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}

type Rectangle struct {
	A Point
	B Point
}

func (rect Rectangle) Range() (float64, float64) {
	return pointRange(rect.A, rect.B)
}

func (rect Rectangle) Center() Point {
	xRange, yRange := rect.Range()
	return Point{X: xRange*0.5 + rect.A.X,
		Y: yRange*0.5 + rect.A.Y}
}

func (rect Rectangle) Radius() float64 {
	c := rect.Center()
	return Line{A: rect.A, B: c}.Distance()
}

func (rect Rectangle) Inside(p Point) bool {
	return rect.A.X <= p.X && p.X <= rect.B.X && rect.A.Y <= p.Y && p.Y <= rect.B.Y
}

type Triangle struct {
	A Point
	B Point
	C Point
}
type Fillable interface {
	Center() Point
	Radius() float64
	Inside(p Point) bool
}
