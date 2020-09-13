package primitives

import (
	"math/rand"

	"github.com/fogleman/ln/ln"
)

type Boxes struct{}

func (b Boxes) Dimensions() (int, int) {
	return int(b.Width()), int(b.Height())
}

func (b Boxes) Width() float64 {
	return 1500.0
}

func (b Boxes) Height() float64 {
	return 1000.0
}

func (b Boxes) Step() float64 {
	return 0.01
}

func (b Boxes) Camera() (fovy, znear, zfar float64) {
	return 50.0, 0.1, 10.0
}

func (b Boxes) Draw(context *ln.Scene, rand *rand.Rand) {
	context.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))
}
