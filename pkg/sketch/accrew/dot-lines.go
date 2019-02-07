package accrew

import (
	"math"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

type pointColor struct {
	Point gg.Point
	Color clr.Color
}

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
	hue := rand.Intn(365)
	margins := rand.Float64() * 300
	spacing := rand.Intn(30) + 10
	delta := rand.Float64()*float64(spacing) + 5
	slog.InfoValues("margins", margins)

	w := dl.Width() - margins*2
	h := dl.Height() - margins*2
	slog.InfoValues("h", h, "w", w)
	x0 := w*rand.Float64() + margins
	y0 := h*rand.Float64() + margins
	x1 := w*rand.Float64() + margins
	y1 := h*rand.Float64() + margins
	x2 := w*rand.Float64() + margins
	y2 := h*rand.Float64() + margins
	slog.InfoValues("rungs", rungs, "delta", delta)
	var pointColors []pointColor

	for i := 0; i < rungs; i++ {
		color := clr.HSV{H: hue, S: (100 / rungs) * i, V: 70}
		x0 += delta
		y0 += delta
		x1 += delta
		y1 += delta
		x2 += delta
		y2 += delta
		points := quadraticBezier(x0, y0, x1, y1, x2, y2, spacing)
		for _, p := range points {
			pointColors = append(pointColors, pointColor{Color: color, Point: p})
		}
	}
	slog.InfoValues("len(point)", len(pointColors))

	x, y := boundingBox(pointColors)
	slog.InfoValues("boundX", x, "boundY", y)

	sw, sh := y.X-x.X, y.Y-x.Y

	slog.InfoValues("sw", sw, "sh", sh)
	//context.DrawRectangle(x.X, x.Y, sw, sh)
	slog.InfoValues("desiredX", (w-sw)/2, "desiredY", (h-sh)/2)
	deltaX := (w-sw)/2 - x.X
	deltaY := (h-sh)/2 - x.Y
	slog.InfoValues("deltaX", deltaX, "deltaY", deltaY)
	context.Translate(deltaX+margins, deltaY+margins)

	for _, p := range pointColors {
		r, g, b := p.Color.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		context.DrawPoint(p.Point.X, p.Point.Y, 2)
		context.Fill()
	}
}

func boundingBox(points []pointColor) (x, y gg.Point) {
	x.X, x.Y = math.MaxFloat64, math.MaxFloat64
	for _, pc := range points {
		if pc.Point.X < x.X {
			x.X = pc.Point.X
		}
		if pc.Point.Y < x.Y {
			x.Y = pc.Point.Y
		}

		if pc.Point.X > y.X {
			y.X = pc.Point.X
		}
		if pc.Point.Y > y.Y {
			y.Y = pc.Point.Y
		}
	}
	return
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
	var result []gg.Point
	for i := 0; i < n; i += sampleRate {
		t := float64(i) / d
		x, y := quadratic(x0, y0, x1, y1, x2, y2, t)
		result = append(result, gg.Point{X: x, Y: y})
	}
	return result
}
