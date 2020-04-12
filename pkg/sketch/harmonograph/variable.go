package harmonograph

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/devinmcgloin/sail/pkg/canvas"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
)

type Single struct{}

func (v Single) Dimensions() (int, int) {
	return 1400, 2000
}

func (v Single) Draw(context *gg.Context, rand *rand.Rand) {
	timeStepSize := 0.0001
	totalCycles := 10.0

	context.SetColor(color.Black)
	context.SetLineWidth(1)

	xFreq := rand.Float64() * 200
	xPhase := rand.Float64() * 200
	xAmp := rand.Float64()*600 + 50
	xDamp := rand.Float64()

	yFreq := rand.Float64() * 200
	yPhase := rand.Float64() * 200
	yAmp := rand.Float64()*600 + 50
	yDamp := rand.Float64()

	slog.InfoValues("xFreq", xFreq,
		"xPhase", xPhase,
		"xAmp", xAmp,
		"xDamp", xDamp)

	slog.InfoValues("yFreq", yFreq,
		"yPhase", yPhase,
		"yAmp", yAmp,
		"yDamp", yDamp,
	)
	context.Translate(canvas.W(context, 0.5), canvas.H(context, 0.5))

	xLast := 0.0
	yLast := 0.0
	for time := 0.0; time < totalCycles; time += timeStepSize {
		x := h(xFreq, xPhase, xAmp, xDamp, time)
		y := h(yFreq, yPhase, yAmp, yDamp, time)
		if xLast != 0.0 && yLast != 0.0 {
			context.DrawLine(xLast, yLast, x, y)
		}
		xLast, yLast = x, y
	}

	context.Stroke()
}

type Dual struct{}

func (v Dual) Dimensions() (int, int) {
	return 1400, 2000
}

func (v Dual) Draw(context *gg.Context, rand *rand.Rand) {
	timeStepSize := 0.0001
	totalCycles := 10.0

	context.SetColor(color.Black)
	context.SetLineWidth(1)

	xFreq := rand.Float64() * 200
	xPhase := rand.Float64() * 200
	xAmp := rand.Float64()*600 + 50
	xDamp := rand.Float64()

	yFreq := rand.Float64() * 200
	yPhase := rand.Float64() * 200
	yAmp := rand.Float64()*600 + 50
	yDamp := rand.Float64()

	slog.InfoValues("xFreq", xFreq,
		"xPhase", xPhase,
		"xAmp", xAmp,
		"xDamp", xDamp)

	slog.InfoValues("yFreq", yFreq,
		"yPhase", yPhase,
		"yAmp", yAmp,
		"yDamp", yDamp,
	)
	context.Translate(canvas.W(context, 0.5), canvas.H(context, 0.5))

	xLast := 0.0
	yLast := 0.0
	for time := 0.0; time < totalCycles; time += timeStepSize {
		x := h(xFreq, xPhase, xAmp, xDamp, time) + h(xFreq, xPhase, xAmp, xDamp, time)
		y := h(yFreq, yPhase, yAmp, yDamp, time) + h(yFreq, yPhase, yAmp, yDamp, time)
		if xLast != 0.0 && yLast != 0.0 {
			context.DrawLine(xLast, yLast, x, y)
		}
		xLast, yLast = x, y
	}

	context.Stroke()
}

type Variable struct{}

func (v Variable) Dimensions() (int, int) {
	return 1400, 2000
}

func (v Variable) Draw(context *gg.Context, rand *rand.Rand) {
	timeStepSize := 0.0001
	totalCycles := 10.0
	order := 5.0

	context.SetColor(color.Black)
	context.SetLineWidth(1)

	xFreq := rand.Float64() * 200
	xPhase := rand.Float64() * 200
	xAmp := rand.Float64()*600 + 50
	xDamp := rand.Float64()

	yFreq := rand.Float64() * 200
	yPhase := rand.Float64() * 200
	yAmp := rand.Float64()*600 + 50
	yDamp := rand.Float64()

	slog.InfoValues("xFreq", xFreq,
		"xPhase", xPhase,
		"xAmp", xAmp,
		"xDamp", xDamp)

	slog.InfoValues("yFreq", yFreq,
		"yPhase", yPhase,
		"yAmp", yAmp,
		"yDamp", yDamp,
	)
	context.Translate(canvas.W(context, 0.5), canvas.H(context, 0.5))

	xLast := 0.0
	yLast := 0.0
	for time := 0.0; time < totalCycles; time += timeStepSize {
		x := h(xFreq, xPhase, xAmp, xDamp, time) * order
		y := h(yFreq, yPhase, yAmp, yDamp, time) * order
		if xLast != 0.0 && yLast != 0.0 {
			context.DrawLine(xLast, yLast, x, y)
		}
		xLast, yLast = x, y
	}

	context.Stroke()
}

type Offset struct{}

func (v Offset) Dimensions() (int, int) {
	return 1400, 2000
}

func (v Offset) Draw(context *gg.Context, rand *rand.Rand) {
	timeStepSize := 0.0001
	totalCycles := 10.0

	context.SetColor(color.Black)
	context.SetLineWidth(1)

	xFreq := []float64{rand.Float64() * 200, rand.Float64() * 200}
	xPhase := []float64{rand.Float64() * 200, rand.Float64() * 200}
	xAmp := []float64{rand.Float64()*600 + 50, rand.Float64()*600 + 50}
	xDamp := []float64{rand.Float64(), rand.Float64()}

	yFreq := []float64{rand.Float64() * 200, rand.Float64() * 200}
	yPhase := []float64{rand.Float64() * 200, rand.Float64() * 200}
	yAmp := []float64{rand.Float64()*600 + 50, rand.Float64()*600 + 50}
	yDamp := []float64{rand.Float64(), rand.Float64()}

	slog.InfoValues("xFreq", xFreq,
		"xPhase", xPhase,
		"xAmp", xAmp,
		"xDamp", xDamp)

	slog.InfoValues("yFreq", yFreq,
		"yPhase", yPhase,
		"yAmp", yAmp,
		"yDamp", yDamp,
	)
	context.Translate(canvas.W(context, 0.5), canvas.H(context, 0.5))

	xLast := 0.0
	yLast := 0.0
	for time := 0.0; time < totalCycles; time += timeStepSize {
		x := h2(xFreq, xPhase, xAmp, xDamp, time)
		y := h2(yFreq, yPhase, yAmp, yDamp, time)
		if xLast != 0.0 && yLast != 0.0 {
			context.DrawLine(xLast, yLast, x, y)
		}
		xLast, yLast = x, y
	}

	context.Stroke()
}

func h2(frequency, phase, amplitude, damping []float64, time float64) float64 {
	termSum := 0.0
	for i := range frequency {
		termSum += h(frequency[i], phase[i], amplitude[i], damping[i], time)
	}
	return termSum
}

func h(frequency, phase, amplitude, damping, time float64) float64 {
	return amplitude * math.Sin(time*frequency+phase) * math.Exp(-1*damping*time)
}
