package library

import (
	"bytes"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
	"github.com/fogleman/ln/ln"
)

type Runner interface {
	Render(seed int64) (*bytes.Buffer, error)
	Dimensions() (width, height int)
	Write(seed int64) error
	Directory() string
	Path(seed int64) string
}

type PNGRunner struct {
	Sketch   sketch.Renderable
	context  *gg.Context
	SketchID string
}

type SVGRunner struct {
	Sketch   sketch.Plotable
	context  *ln.Scene
	SketchID string
}

// Render renders the given sketch to the filessystem with the provided seed
func (pngRender *PNGRunner) Render(seed int64) (*bytes.Buffer, error) {
	xd, yd := pngRender.Dimensions()
	context := gg.NewContext(xd, yd)
	pngRender.context = context

	slog.InfoPrintf("Rendering %T with dimesions (%d, %d) and seed %d\n", pngRender.Sketch, xd, yd, seed)
	rand := rand.New(rand.NewSource(seed))
	clearBackground(context)

	pngRender.Sketch.Draw(context, rand)
	bytes := new(bytes.Buffer)

	err := context.EncodePNG(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (pngRunner *PNGRunner) Write(seed int64) error {
	path := pngRunner.Path(seed)
	return pngRunner.context.SavePNG(path)
}

func (pngRunner *PNGRunner) Dimensions() (width, height int) {
	return pngRunner.Sketch.Dimensions()
}

func (p *PNGRunner) Directory() string {
	return fmt.Sprintf("sketches/%s/", p.SketchID)
}

func (p *PNGRunner) Path(seed int64) string {
	return fmt.Sprintf("sketches/%s/sketch-%d.png", p.SketchID, seed)
}

func clearBackground(context *gg.Context) {
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, float64(context.Width()), float64(context.Height()))
	context.Fill()
}

// func (svgRunner SVGRunner) Dimensions() (width, height int) {
// 	return svgRunner.sketch.Dimensions()
// }

// // Render renders the given sketch to the filessystem with the provided seed
// func (svgRunner SVGRunner) Render(seed int64) (*bytes.Buffer, error) {
// 	xd, yd := svgRunner.Dimensions()
// 	context := gg.NewContext(xd, yd)

// 	slog.InfoPrintf("Rendering %T with dimesions (%d, %d) and seed %d\n", pngRender.sketch, xd, yd, seed)
// 	rand := rand.New(rand.NewSource(seed))
// 	clearBackground(context)

// 	pngRender.sketch.Draw(context, rand)
// 	bytes := new(bytes.Buffer)

// 	err := context.EncodePNG(bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes, nil
// }
