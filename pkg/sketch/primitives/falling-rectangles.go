package primitives

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

//FallingRectangles defines the sketch type
type FallingRectangles struct{}

// Dimensions returns the size of the sketch in integers
func (fr FallingRectangles) Dimensions() (int, int) {
	return int(fr.Width()), int(fr.Height())
}

// Width helper for finding the width of the canvas
func (fr FallingRectangles) Width() float64 {
	return 800.0
}

// Height helper for finding the hieght of the canvas
func (fr FallingRectangles) Height() float64 {
	return 1300.0
}

// Draw is the primary rendering method
func (fr FallingRectangles) Draw(context *gg.Context, rand *rand.Rand) {
	margin := rand.Float64()*200 + 10

	avaliableSpace := fr.Width() - margin*2
	sizeFactor := math.Floor(rand.Float64()*25 + 5)
	noiseFactor := rand.Float64() * 2
	boxSize := avaliableSpace / sizeFactor
	halfBox := boxSize / 2

	context.SetLineWidth(1)
	context.SetColor(color.Black)

	fmt.Printf("\tMargin: %f\n\tAvaliableSpace: %f\n\tsizeFactor: %f\n\tboxSize: %f\n", margin, avaliableSpace, sizeFactor, boxSize)

	for x := margin + halfBox; x < fr.Width()-margin; x += boxSize {
		rowIndex := 0.0
		for y := margin + halfBox; y < fr.Height()-margin; y += boxSize {
			normalizedNoise := (rowIndex / (fr.Width() - margin)) * (noiseFactor * 10)
			rotate := normalizedNoise * rand.NormFloat64()
			context.Push()
			context.RotateAbout(rotate, x, y)
			context.Translate(rand.Float64()*rowIndex*noiseFactor, rand.Float64()*rowIndex*noiseFactor)
			context.DrawRectangle(x-halfBox, y-halfBox, boxSize, boxSize)
			context.Stroke()
			context.Pop()
			rowIndex++
		}
	}
}

func drawAbsRect(dc *gg.Context, x1, y1, x2, y2 float64) {
	dc.NewSubPath()
	dc.MoveTo(x1, y1)
	dc.LineTo(x2, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x1, y2)
	dc.ClosePath()
}
