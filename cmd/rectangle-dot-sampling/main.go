package main

import (
	"fmt"
	"image/color"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/generative-go/pkg/canvas"
	"github.com/devinmcgloin/generative-go/pkg/fill"
	"github.com/devinmcgloin/generative-go/pkg/shapes"
	"github.com/fogleman/gg"
)

func main() {
	context := gg.NewContext(1400, 900)
	var rows = 10.0
	var margin = 0.05
	filler := fill.NewUniformFiller(8000, 42)
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

		r, g, b := clr.HSV{H: 34, S: uint8(i * 7), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))

		fmt.Printf("%+v R: %d, G: %d, B: %d\n", rect, r, g, b)
		filler.DotFill(context, rect)
	}
	context.SavePNG("out.png")
}

func rectangePositioning(offset, index, rectangeCount float64) float64 {
	avaliableSpace := 1.0 - offset*2
	return offset + index*avaliableSpace/rectangeCount
}
