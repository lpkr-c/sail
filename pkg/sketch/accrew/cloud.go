package accrew

import (
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

type Cloud struct{}

func (c Cloud) Dimensions() (int, int) {
	return 800, 1400
}

func (c Cloud) Draw(context *gg.Context, rand *rand.Rand) {
	rows := rand.Float64() * 200
	hue := rand.Intn(365)
	growthFactor := rand.Float64() * 50
	minGrowth := rand.Float64() * 20

	slog.InfoValues("rows", rows, "hue", hue, "growth-factor", growthFactor, "min-growth", minGrowth)
	accrew := make([]float64, 100)
	for i := range accrew {
		accrew[i] = 100
	}

	for i := 0.0; i < rows; i++ {
		r, g, b := clr.HSV{H: hue, S: int(i), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		for i, height := range accrew {
			incBy := rand.Float64()*float64(rows)/growthFactor + minGrowth
			context.DrawPoint(float64(i*6)+100, incBy+height, 2)
			context.Fill()
			accrew[i] += incBy
		}
	}

}
