package primitives

import (
	"image/color"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/fogleman/gg"
)

type LineColoring struct{}

func (lc LineColoring) Dimensions() (int, int) {
	return 1000, 1000
}

func (lc LineColoring) Draw(context *gg.Context, rand *rand.Rand) {
	rows := rand.Float64() * 200
	margin := rand.Float64()*200 + 10
	noiseFactor := rand.Float64() * 10
	hue := uint16(rand.Intn(365))

	avaliableSpace := 1000 - margin*2
	spacing := avaliableSpace / rows

	context.SetColor(color.Black)
	context.SetLineWidth(3)

	for x := margin; x < 1000-margin; x += spacing {
		for y := margin; y < 1000-margin; y += spacing {
			n := (rand.Float64()*2 - 1) * noiseFactor
			r, g, b := clr.HSV{H: hue, S: uint8(n), V: 70}.RGB()
			context.SetRGB(float64(r), float64(g), float64(b))
			context.DrawLine(x+n, y+n, x-n, y-n)
			context.Stroke()
		}
	}
}
