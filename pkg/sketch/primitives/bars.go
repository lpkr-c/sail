package primitives

import (
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/fogleman/gg"
)

type Bars struct{}

func (b Bars) Dimensions() (int, int) {
	return int(b.Width()), int(b.Height())
}

func (b Bars) Width() float64 {
	return 1500.0
}
func (b Bars) Height() float64 {
	return 1000.0
}

func (b Bars) Draw(context *gg.Context, rand *rand.Rand) {
	rows := rand.Float64()*200 + 50
	margin := rand.Float64()*200 + 10
	noiseFactor := rand.Float64() * 10
	hue := uint16(rand.Intn(365))
	lineWidth := rand.Int31n(10)

	avaliableSpace := b.Width() - margin*2
	spacing := avaliableSpace / rows

	context.SetLineWidth(float64(lineWidth))

	y1 := margin
	y2 := b.Height() - margin
	for x := margin; x < b.Width()-margin; x += spacing {
		n := (rand.Float64()*2 - 1) * noiseFactor
		r, g, b := clr.HSV{H: hue, S: uint8(n), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		context.DrawLine(x+n, y1, x+n, y2)
		context.Stroke()
	}
}
