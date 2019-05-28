package gradients

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

// Skyspace defines the type of the sketch
type Skyspace struct {
}

// Dimensions determines how large it should be
func (ss Skyspace) Dimensions() (int, int) {
	return 2000, 1400
}

// Draw actually renders the sketch onto the canvas
func (ss Skyspace) Draw(context *gg.Context, rand *rand.Rand) {
	hueVariance := rand.Intn(85)
	hue := rand.Intn(365)
	radialHue := hue + (rand.Intn(hueVariance) - hueVariance/2)
	slog.InfoValues("hueVariance", hueVariance, "hue", hue, "radialHue", radialHue)

	center := clr.HSV{H: hue, S: 100, V: 100}
	surroundingColor := clr.HSV{H: radialHue, S: 85, V: 20}
	slog.InfoValues("center", center, "surroundingcolor", surroundingColor)

	rectGradient := SquareDistanceGradient{
		from: center, to: surroundingColor,
		x1: 300, y1: 300, x2: 1700, y2: 1100, maxDistance: 2000,
	}
	slog.InfoValues("dist", rectGradient.Distance(0, 0))
	slog.InfoValues("dist", rectGradient.Distance(500, 500))

	context.SetFillStyle(rectGradient)
	context.DrawRectangle(0, 0, 2000, 2000)
	context.Fill()

	linearGradient := gg.NewLinearGradient(300, 300, 1400, 800)
	linearGradient.AddColorStop(0, center)
	linearGradient.AddColorStop(1, center)

	context.SetFillStyle(linearGradient)
	context.DrawRoundedRectangle(300, 300, 1400, 800, 4)
	// context.DrawRoundedRectangle(0, 0, 100, 100, 4)

	context.Fill()
}

type SquareDistanceGradient struct {
	from           clr.HSV
	to             clr.HSV
	x1, y1, x2, y2 int
	maxDistance    float64
}

func (sdg SquareDistanceGradient) ColorAt(x, y int) color.Color {
	dist := sdg.Distance(x, y) / sdg.maxDistance
	return lerp(sdg.from, sdg.to, dist)
}

func (sgd SquareDistanceGradient) Distance(x, y int) float64 {
	dx := max(sgd.x1-x, 0, x-sgd.x2)
	dy := max(sgd.y1-y, 0, y-sgd.y2)
	// slog.InfoValues("x", x, "y", y, "dx", dx, "dy", dy)
	return math.Sqrt(dx*dx + dy*dy)
}

func lerp(a, b clr.HSV, t float64) clr.HSV {
	var h, s, v float64

	var h1, s1, v1 = float64(a.H) / 365, float64(a.S) / 100, float64(a.V) / 100
	var h2, s2, v2 = float64(b.H) / 365, float64(b.S) / 100, float64(b.V) / 100

	d := h2 - h1
	if h1 > h2 {
		h1, h2 = h2, h1

		d = -d
		t = 1 - t
	}

	if d > 0.5 {
		h1 = h1 + 1
		h = math.Mod(h1+t*(h2-h1), 1.0)
	} else if d <= 0.5 {
		h = h1 + t*d
	}

	s = s1 + t*(s2-s1)
	v = v1 + t*(v2-v1)

	return clr.HSV{
		H: int(h * 365),
		S: int(s * 100),
		V: int(v * 100),
	}
}

func max(args ...int) float64 {
	max := 0.0
	for _, val := range args {
		if float64(val) > max {
			max = float64(val)
		}
	}
	return max
}
