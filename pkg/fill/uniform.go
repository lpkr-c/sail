package fill

import (
	"math/rand"

	"github.com/devinmcgloin/generative-go/pkg/shapes"
	"github.com/fogleman/gg"
)

type DotFiller struct {
	N    int
	Rand *rand.Rand
}

func (f DotFiller) Float() float64 {
	return f.Rand.Float64()
}

func (f DotFiller) FillRectangle(canvas *gg.Context, rect shapes.Rectangle) {
	for i := 0; i < f.N; i++ {
		var xRange, yRange = rect.Range()
		var xDiff = f.Float() * float64(xRange)
		var yDiff = f.Float() * float64(yRange)
		canvas.DrawPoint(rect.A.X+xDiff, rect.A.Y+yDiff, 1)
		canvas.Fill()
	}
}
