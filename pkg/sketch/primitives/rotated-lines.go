package primitives

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/fogleman/gg"
)

type RotatedLines struct{}

func (rl RotatedLines) Dimensions() (int, int) {
	return 1000, 1000
}

func (rl RotatedLines) Draw(context *gg.Context, rand *rand.Rand) {
	rows := rand.Float64() * 200
	margin := rand.Float64() * 200

	avaliableSpace := 1000 - margin*2
	spacing := avaliableSpace / rows
	context.SetColor(color.Black)
	context.SetLineWidth(5)

	fmt.Printf("Rows: %f Margin: %f Spacing: %f\n", rows, margin, spacing)
	for x := margin; x < avaliableSpace; x += spacing {
		for y := margin; y < avaliableSpace; y += spacing {
			n := rand.Float64() * 15
			r, g, b := clr.HSV{H: uint16(x), S: uint8(x + y), V: 70}.RGB()
			context.SetRGB(float64(r), float64(g), float64(b))
			fmt.Printf("X1: %f Y1: %f X2: %f Y2: %f\n", x+n, y+n, x-n, y-n)
			context.DrawLine(x+n, y+n, x-n, y-n)
			context.Stroke()
		}
	}
}
