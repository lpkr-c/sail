package topography

import (
	"math/rand"

	"github.com/fogleman/gg"
)

type HillClimbing struct {
}

func (hc HillClimbing) Dimensions() (int, int) {
	return 1400, 900
}

func (hc HillClimbing) Draw(context *gg.Context, r *rand.Rand) {
	// noise := noise.New(rand.Int63())
}

func gradientDescent() {}
