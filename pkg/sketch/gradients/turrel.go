package gradients

import (
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/fogleman/gg"
)

// Skyspace defines the type of the sketch
type Skyspace struct {
}

// Dimensions determines how large it should be
func (ss Skyspace) Dimensions() (int, int) {
	return 1400, 1400
}

// Draw actually renders the sketch onto the canvas
func (ss Skyspace) Draw(context *gg.Context, rand *rand.Rand) {
	surroundingColorA := clr.RGB{
		R: 255, G: 255, B: 255,
	}
	surroundingColorB := clr.RGB{
		R: 216, G: 88, B: 22,
	}

	radialGradient := gg.NewRadialGradient(700, 700, 200, 700, 700, 1000)
	radialGradient.AddColorStop(0, surroundingColorA)
	radialGradient.AddColorStop(1, surroundingColorB)
	context.SetFillStyle(radialGradient)
	context.DrawRectangle(0, 0, 1400, 1400)
	context.Fill()

	centerColorA := clr.RGB{R: 46, G: 66, B: 130}
	centerColorB := clr.RGB{R: 93, G: 120, B: 176}
	linearGradient := gg.NewLinearGradient(300, 300, 800, 800)
	linearGradient.AddColorStop(0, centerColorA)
	linearGradient.AddColorStop(1, centerColorB)

	context.SetFillStyle(linearGradient)
	context.DrawRectangle(300, 300, 800, 800)
	context.Fill()

}
