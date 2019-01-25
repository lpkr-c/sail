package accrew

import (
	"math"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

//DotLines defines the type of sketch
type DotLines struct{}

//Dimensions determes the size of the sketch to be rendered
func (dl DotLines) Dimensions() (int, int) {
	return int(dl.Width()), int(dl.Height())
}

// Width defines the width of the sketch
func (dl DotLines) Width() float64 {
	return 1000.0
}

// Height defines the width of the sketch
func (dl DotLines) Height() float64 {
	return 1000.0
}

// Draw renders the sketch on the given context with the seed
func (dl DotLines) Draw(context *gg.Context, rand *rand.Rand) {
	rungs := rand.Intn(30) + 5
	hue := uint16(rand.Intn(365))
	margins := rand.Float64() * 300
	spacing := rand.Intn(30) + 10
	delta := rand.Float64()*float64(spacing) + 5

	w := dl.Width() - margins*2
	h := dl.Height() - margins*2
	x0 := w*rand.Float64() + margins
	y0 := h*rand.Float64() + margins
	x1 := w*rand.Float64() + margins
	y1 := h*rand.Float64() + margins
	x2 := w*rand.Float64() + margins
	y2 := h*rand.Float64() + margins
	slog.InfoValues("rungs", rungs, "delta", delta)
	for i := 0; i < rungs; i++ {
		r, g, b := clr.HSV{H: hue, S: uint8((100 / rungs) * i), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		x0 += delta
		y0 += delta
		x1 += delta
		y1 += delta
		x2 += delta
		y2 += delta
		points := quadraticBezier(x0, y0, x1, y1, x2, y2, spacing)
		for _, p := range points {
			context.DrawPoint(p.X, p.Y, 2)
			context.Fill()
		}
	}
}

func quadratic(x0, y0, x1, y1, x2, y2, t float64) (x, y float64) {
	u := 1 - t
	a := u * u
	b := 2 * u * t
	c := t * t
	x = a*x0 + b*x1 + c*x2
	y = a*y0 + b*y1 + c*y2
	return
}

func quadraticBezier(x0, y0, x1, y1, x2, y2 float64, sampleRate int) []gg.Point {
	l := (math.Hypot(x1-x0, y1-y0) +
		math.Hypot(x2-x1, y2-y1))
	n := int(l + 0.5)
	if n < 4 {
		n = 4
	}
	d := float64(n) - 1
	result := make([]gg.Point, n)
	for i := 0; i < n; i += sampleRate {
		t := float64(i) / d
		x, y := quadratic(x0, y0, x1, y1, x2, y2, t)
		result[i] = gg.Point{X: x, Y: y}
	}
	return result
}
