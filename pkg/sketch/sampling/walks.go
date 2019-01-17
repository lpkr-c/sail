package sampling

import (
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

type DotWalk struct{}

func (dw DotWalk) Dimensions() (int, int) {
	return 1000, 1000
}

func (dw DotWalk) Draw(context *gg.Context, rand *rand.Rand) {
	walkers := rand.Intn(1000)
	hue := uint16(rand.Intn(365))
	steps := rand.Intn(500)
	radius := rand.Float64() + 1
	alpha := rand.Float64()*5 + 1
	margin := rand.Float64()*200 + 10

	slog.InfoValues("walkers", walkers, "hue", hue, "steps", steps, "radius", radius, "alpha", alpha)

	for walker := 0; walker < walkers; walker++ {
		r, g, b := clr.HSV{H: hue, S: uint8(rand.Intn(walkers) * 10), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		xLimit, yLimit := dw.Dimensions()
		x := rand.Float64()*(float64(xLimit)-margin*2) + margin
		y := rand.Float64()*(float64(yLimit)-margin*2) + margin

		for step := 0; step < steps; step++ {
			context.DrawCircle(x, y, radius)
			context.Stroke()

			xDelta := rand.NormFloat64() * alpha
			yDelta := rand.NormFloat64() * alpha
			if Outside(margin, xDelta+x, float64(xLimit)) {
				xDelta *= -1.0
			}

			if Outside(margin, yDelta+y, float64(yLimit)) {
				yDelta *= -1.0
			}

			x += xDelta
			y += yDelta
		}
	}
}

func Outside(margin, position, limit float64) bool {
	return position <= margin || position >= (limit-margin)
}
