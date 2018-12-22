package fill

import (
	"math"
	"math/rand"

	"github.com/devinmcgloin/sail/pkg/shapes"
	"github.com/fogleman/gg"
)

type DotFiller interface {
	Fill(context *gg.Context, shape shapes.Fillable)
}

// RadialFiller represents the basic properties needed to execute various types of fills.
type RadialFiller struct {
	N    int
	rand *rand.Rand
}

// NewRadialFiller sets up the fields needed to crete a Radial Filler
func NewRadialFiller(n int, r *rand.Rand) *RadialFiller {
	return &RadialFiller{
		N:    n,
		rand: r,
	}
}

// DotFill takes any shape that implements Fillable and samples random points inside the shape
func (f RadialFiller) Fill(canvas *gg.Context, shape shapes.Fillable) {
	pointsPlaced := 0
	for pointsPlaced < f.N {
		center := shape.Center()
		radius := shape.Radius()

		r := f.rand.NormFloat64() * radius
		theta := f.rand.NormFloat64() * radius
		xDiff := r * math.Cos(theta)
		yDiff := r * math.Sin(theta)
		p := shapes.Point{X: center.X + xDiff, Y: center.Y + yDiff}
		if shape.Inside(p) {
			pointsPlaced++
			canvas.DrawPoint(p.X, p.Y, 1)
			canvas.Fill()
		}
	}
}

// UniformFiller represents the basic properties needed to execute various types of fills.
type UniformFiller struct {
	N    int
	rand *rand.Rand
}

// NewUniformFiller sets up the fields needed to crete a UniformFiller
func NewUniformFiller(n int, r *rand.Rand) *UniformFiller {
	return &UniformFiller{
		N:    n,
		rand: r,
	}
}

// DotFill takes any shape that implements Fillable and samples random points inside the shape
func (f UniformFiller) DotFill(canvas *gg.Context, shape shapes.Fillable) {
	pointsPlaced := 0
	for pointsPlaced < f.N {
		center := shape.Center()
		radius := shape.Radius()

		r := (f.rand.Float64() * radius) * (1 / 2 * math.Pi)
		theta := f.rand.NormFloat64() * radius
		xDiff := r * math.Cos(theta)
		yDiff := r * math.Sin(theta)
		p := shapes.Point{X: center.X + xDiff, Y: center.Y + yDiff}
		if shape.Inside(p) {
			pointsPlaced++
			canvas.DrawPoint(p.X, p.Y, 1)
			canvas.Fill()
		}
	}
}
