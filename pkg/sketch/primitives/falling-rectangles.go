package primitives

import (
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

func drawRectangle(c *gg.Context, x, y float64, noise float64, size float64) {
}

type FallingRectangles struct{}

func (rg FallingRectangles) Dimensions() (int, int) {
	return int(rg.Width()), int(rg.Height())
}

func (rg FallingRectangles) Width() float64 {
	return 800.0
}
func (rg FallingRectangles) Height() float64 {
	return 1300.0
}

func (fr FallingRectangles) Draw(context *gg.Context, rand *rand.Rand) {
	margin := rand.Float64()*200 + 10

	avaliableSpace := fr.Width() - margin*2
	columns := rand.Float64()*25 + 5
	noiseFactor := rand.Float64()*200 + 1
	boxSize := avaliableSpace / columns

	context.SetColor(color.Black)

	for x := margin; x < fr.Width()-margin; x += boxSize {
		rowIndex := 0.0
		for y := margin; y < fr.Height()-margin; y += boxSize {
			rotated := rand.Float64() * rowIndex / noiseFactor
			context.Push()
			context.RotateAbout(rotated, x, y)
			context.DrawRectangle(x, y, boxSize, boxSize)
			context.Stroke()
			context.Pop()
			rowIndex++
		}
	}
}
