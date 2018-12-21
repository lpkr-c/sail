package sampling

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/canvas"
	"github.com/devinmcgloin/sail/pkg/fill"
	"github.com/devinmcgloin/sail/pkg/shapes"
	"github.com/fogleman/gg"
)

type RectangleDot struct{}

func (c RectangleDot) Dimensions() (int, int) {
	return 1400, 900
}

func (c RectangleDot) Draw(context *gg.Context, r *rand.Rand) {
	rows := 1 + math.Floor(r.Float64()*15)
	margin := r.Float64() * 0.10
	hue := uint16(r.Intn(365))

	filler := fill.NewUniformFiller(8000, r)
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, 1400, 900)
	context.Fill()

	for i := 0.0; i < rows; i++ {
		rect := shapes.Rectangle{
			A: shapes.Point{
				X: canvas.W(context, rectangePositioning(margin, i, rows)),
				Y: canvas.H(context, margin),
			},
			B: shapes.Point{
				X: canvas.W(context, rectangePositioning(margin, 1+i, rows)),
				Y: canvas.H(context, 1-margin),
			},
		}

		r, g, b := clr.HSV{H: hue, S: uint8(i * 7), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))

		filler.DotFill(context, rect)
	}
}

func rectangePositioning(offset, index, rectangeCount float64) float64 {
	avaliableSpace := 1.0 - offset*2
	return offset + index*(avaliableSpace/rectangeCount)
}
