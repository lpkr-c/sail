package sketch

import (
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/fogleman/ln/ln"
)

type Renderable interface {
	Draw(context *gg.Context, rand *rand.Rand)
	Dimensions() (width, height int)
}

type Plotable interface {
	Draw(context *ln.Scene, rand *rand.Rand)
	Dimensions() (width, height int)
	Step() float64
	Camera() (fovy, znear, zfar float64)
}
