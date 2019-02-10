package accrew

import (
	"math/rand"

	"github.com/devinmcgloin/clr/clr"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

type DotCloud struct{}

func (c DotCloud) Dimensions() (int, int) {
	return 800, 1400
}

func (c DotCloud) Draw(context *gg.Context, rand *rand.Rand) {
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

type DisjointLineCloud struct{}

func (c DisjointLineCloud) Dimensions() (int, int) {
	return 800, 1400
}

func (c DisjointLineCloud) Draw(context *gg.Context, rand *rand.Rand) {
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
		for i := range accrew {
			incBy := rand.Float64()*float64(rows)/growthFactor + minGrowth
			accrew[i] += incBy
		}

		for i, height := range accrew {
			x1, y1 := float64(i*6)+100, height
			x2, y2 := float64((i+1)*6)+100, height
			x3, y3 := float64((i+2)*6)+100, height
			context.CubicTo(x1, y1, x2, y2, x3, y3)
			context.Stroke()
		}
	}

}

type JointLineCloud struct{}

func (c JointLineCloud) Dimensions() (int, int) {
	return 800, 1400
}

func (c JointLineCloud) Draw(context *gg.Context, rand *rand.Rand) {
	rows := rand.Float64() * 400
	hue := rand.Intn(365)
	minGrowth := rand.Float64() * 5
	growthFactor := rand.Float64()*50 + 20

	slog.InfoValues("rows", rows, "hue", hue, "growth-factor", growthFactor, "min-growth", minGrowth)
	accrew := make([]float64, 100)
	for i := range accrew {
		accrew[i] = 100
	}

	deltas := make([]float64, 100)

	for i := 0.0; i < rows; i++ {
		r, g, b := clr.HSV{H: hue, S: int(i), V: 70}.RGB()
		context.SetRGB(float64(r), float64(g), float64(b))
		for i := range deltas {
			deltas[i] = rand.Float64()*float64(rows)/growthFactor + minGrowth
		}

		accrew[0] += avg(deltas[0], deltas[1], deltas[2])
		for i := 1; i < len(deltas)-1; i++ {
			deltas[i] = avg(deltas[i-1], deltas[i], deltas[i+1])
			accrew[i] += deltas[i]
		}
		accrew[len(deltas)-1] += avg(deltas[len(deltas)-3], deltas[len(deltas)-1], deltas[len(deltas)-2])

		for i := 0; i < len(deltas); i++ {
			deltas[i] = 0
		}

		for i := 0; i < len(accrew)-1; i++ {
			x1, y1 := float64(i*6)+100, accrew[i]
			x2, y2 := float64((i+1)*6)+100, accrew[i+1]
			context.QuadraticTo(x1, y1, x2, y2)
			context.Stroke()
		}
	}
}

func avg(items ...float64) float64 {
	sum := 0.0
	for _, item := range items {
		sum += item
	}
	return sum / float64(len(items))
}
