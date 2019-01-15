package sketch

import (
	"math/rand"

	"github.com/fogleman/gg"
)

type Renderable interface {
	Draw(context *gg.Context, rand *rand.Rand)
	Dimensions() (int, int)
	// AspectRatio() float64
}
