package main

import (
	"image/color"

	"github.com/devinmcgloin/clr/clr"
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

	tri := shapes.Triangle{
		A: shapes.Point{},
		B: shapes.Point{},
		C: shapes.Point{},
	}

	r, g, b := clr.HSV{H: 34, S: 67, V: 70}.RGB()
	context.SetRGB(float64(r), float64(g), float64(b))

	filler.DotFill(context, tri)

	context.SavePNG("out.png")
}

func rectangePositioning(offset, index, rectangeCount float64) float64 {
	avaliableSpace := 1.0 - offset*2
	return offset + index*avaliableSpace/rectangeCount
}
