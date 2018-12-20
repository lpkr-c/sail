package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/devinmcgloin/clr/clr"
	"github.com/fogleman/gg"
)

func main() {
	seed := time.Now().Unix()
	rand := rand.New(rand.NewSource(seed))

	context := gg.NewContext(800, 1400)
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, 800, 1400)
	context.Fill()

	rows := rand.Float64() * 200
	hue := uint16(rand.Intn(365))
	growthFactor := rand.Float64() * 50
	minGrowth := rand.Float64() * 20

	accrew := make([]float64, 100)
	for i := range accrew {
		accrew[i] = 100
	}

	for i := 0.0; i < rows; i++ {
		r, g, b := clr.HSV{H: hue, S: uint8(i), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		for i, height := range accrew {
			incBy := rand.Float64()*float64(rows)/growthFactor + minGrowth
			context.DrawPoint(float64(i*6)+100, incBy+height, 2)
			context.Fill()
			accrew[i] += incBy
		}
	}
	err := context.SavePNG(fmt.Sprintf("./sketches/cloud-lines/%d-sketch.png", seed))
	if err != nil {
		fmt.Println(err)
	}
}
