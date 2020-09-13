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
	paths    ln.Paths
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

// Render renders the given sketch to the filessystem with the provided seed
func (s *SVGRunner) Render(seed int64) (*bytes.Buffer, error) {
	xd, yd := s.Dimensions()
	context := &ln.Scene{}

	slog.InfoPrintf("Rendering %T with dimesions (%d, %d) and seed %d\n", s.Sketch, xd, yd, seed)
	rand := rand.New(rand.NewSource(seed))

	s.Sketch.Draw(context, rand)

	eye := ln.Vector{4, 3, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}

	fovy := 50.0 // vertical field of view, degrees
	znear := 0.1 // near z plane
	zfar := 10.0 // far z plane
	step := s.Sketch.Step()

	paths := context.Render(eye, center, up, float64(xd), float64(yd), fovy, znear, zfar, step)
	s.paths = paths
	bytes := new(bytes.Buffer)

	svg := paths.ToSVG(float64(xd), float64(yd))

	_, err := bytes.WriteString(svg)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *SVGRunner) Write(seed int64) error {
	path := s.Path(seed)
	xd, yd := s.Dimensions()

	return s.paths.WriteToSVG(path, float64(xd), float64(yd))
}

func (s *SVGRunner) Dimensions() (width, height int) {
	return s.Sketch.Dimensions()
}

func (s *SVGRunner) Directory() string {
	return fmt.Sprintf("sketches/%s/", s.SketchID)
}

func (s *SVGRunner) Path(seed int64) string {
	return fmt.Sprintf("sketches/%s/sketch-%d.svg", s.SketchID, seed)
}
